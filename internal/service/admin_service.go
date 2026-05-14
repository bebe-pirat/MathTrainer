package service

import (
	"MathTrainer/internal"
	"MathTrainer/internal/model"
	"MathTrainer/internal/repository"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AdminService interface {
	CreateSchool(ctx context.Context, name string, address string) error
	GetSchools(ctx context.Context) ([]model.School, error)
	GetTeachersBySchoolId(ctx context.Context, schoolId int) ([]model.User, error)

	CreateTeacher(ctx context.Context, fullName, login, email string, classId int) (*model.UserCredentials, error)
	ChangeBlockingUser(ctx context.Context, userId int, blocked bool) error
	GetAllUsers(ctx context.Context) ([]model.User, error)
	CreateUser(ctx context.Context, e model.CreateAndUpdateUserRequest) (*model.UserCredentials, error)
	UpdateUser(ctx context.Context, e model.CreateAndUpdateUserRequest) error
	DeleteUser(ctx context.Context, id int) error
	UpdatePassword(ctx context.Context, userId int) (*model.UserCredentials, error)

	CreateSection(ctx context.Context, section model.Section) error
	UpdateSection(ctx context.Context, section model.Section) error
	DeleteSection(ctx context.Context, sectionId int) error
	GetSections(ctx context.Context, class int) ([]model.Section, error)

	CreateEquationType(ctx context.Context, e model.CreateEquationTypeRequest) error
	DeleteEquationType(ctx context.Context, equationId int) error
	UpdateEquationType(ctx context.Context, equationTypeId int, e model.UpdateEquationTypeRequest) error
	GetEquationTypes(ctx context.Context) ([]model.EquationType, error)
	GetOperandsForEquationType(ctx context.Context, equationTypeId int) ([]model.OperandResponse, error)

	JoinEquationTypeAndSection(ctx context.Context, equationTypeId, sectionId int) error
	UnJoinEquationTypeAndSection(ctx context.Context, equationTypeId, sectionId int) error
	GetSectionsAndEquationTypes(ctx context.Context) ([]model.SectionAndEquationType, error)
}

type AdminServiceStruct struct {
	userRepo         repository.UserRepository
	schoolRepo       repository.SchoolRepository
	sectionRepo      repository.SectionRepository
	equationTypeRepo repository.EquationTypeRepository
}

func NewAdminServiceStruct(userRepo repository.UserRepository, schoolRepo repository.SchoolRepository, sectionRepo repository.SectionRepository, equationTypeRepo repository.EquationTypeRepository) *AdminServiceStruct {
	return &AdminServiceStruct{
		userRepo:         userRepo,
		schoolRepo:       schoolRepo,
		sectionRepo:      sectionRepo,
		equationTypeRepo: equationTypeRepo,
	}
}

func (s *AdminServiceStruct) CreateSchool(ctx context.Context, name string, address string) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return errors.New("bad request")
	}

	address = strings.TrimSpace(address)
	if address == "" {
		return errors.New("bad request")
	}

	school := model.School{
		Name:       name,
		Address:    address,
		Created_at: time.Now(),
	}

	return s.schoolRepo.CreateSchool(ctx, school)
}

func (s *AdminServiceStruct) GetSchools(ctx context.Context) ([]model.School, error) {
	return s.schoolRepo.GetAllSchools(ctx)
}

func (s *AdminServiceStruct) GetTeachersBySchoolId(ctx context.Context, schoolId int) ([]model.User, error) {
	return s.userRepo.GetTeachersBySchool(ctx, schoolId)
}

func (s *AdminServiceStruct) CreateTeacher(ctx context.Context, fullName, login, email string, classId int) (*model.UserCredentials, error) {
	slog.Info("parameters", "fullname", fullName, "email", email, "login", login, "class_id", classId)

	fullName = strings.TrimSpace(fullName)
	if fullName == "" {
		return nil, model.BadRequest("no fullname")
	}

	email = strings.TrimSpace(email)
	if email == "" {
		return nil, model.BadRequest("no email")
	}

	login = strings.TrimSpace(login)
	if login == "" {
		return nil, model.BadRequest("no login")
	}

	password, err := GenerateRandomPassword()
	if err != nil {
		return nil, err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	teacher := model.User{
		Email:        email,
		Login:        login,
		FullName:     fullName,
		PasswordHash: passwordHash,
		RoleId:       internal.TeacherRoleId,
		Blocked:      false,
		ClassId:      &classId,
		CreatedAt:    time.Now(),
	}

	login, err = s.userRepo.CreateUser(ctx, teacher)
	if err != nil {
		return nil, err
	}

	return &model.UserCredentials{Login: login, Password: password}, nil
}

func (s *AdminServiceStruct) ChangeBlockingUser(ctx context.Context, userId int, blocked bool) error {
	if userId <= 0 {
		return errors.New("bad request")
	}

	return s.userRepo.BlockUser(ctx, userId, blocked)
}

func (s *AdminServiceStruct) GetAllUsers(ctx context.Context) ([]model.User, error) {
	return s.userRepo.GetAllUsers(ctx)
}

func (s *AdminServiceStruct) CreateSection(ctx context.Context, section model.Section) error {
	section.Name = strings.TrimSpace(section.Name)
	if section.Name == "" || section.Class < 1 || section.Class > 4 {
		return model.ErrBadRequest
	}

	return s.sectionRepo.CreateSection(ctx, section)
}

func (s *AdminServiceStruct) UpdateSection(ctx context.Context, section model.Section) error {
	section.Name = strings.TrimSpace(section.Name)
	if section.Id <= 0 || section.Name == "" || section.Class < 1 || section.Class > 4 || section.Order < 0 {
		return model.ErrBadRequest
	}

	err := s.sectionRepo.UpdateSection(ctx, section)
	if err == sql.ErrNoRows {
		return model.ErrNotFound
	}

	return err
}

func (s *AdminServiceStruct) DeleteSection(ctx context.Context, sectionId int) error {
	if sectionId <= 0 {
		return model.ErrBadRequest
	}

	err := s.sectionRepo.DeleteSection(ctx, sectionId)
	if err == sql.ErrNoRows {
		return model.ErrNotFound
	}

	return err
}

func (s *AdminServiceStruct) GetSections(ctx context.Context, class int) ([]model.Section, error) {
	return s.sectionRepo.GetSections(ctx, class)
}

func (s *AdminServiceStruct) CreateUser(ctx context.Context, e model.CreateAndUpdateUserRequest) (*model.UserCredentials, error) {
	e.Email = strings.TrimSpace(e.Email)
	if e.Email == "" {
		return nil, model.BadRequest("no EMAIL")
	}

	e.Login = strings.TrimSpace(e.Login)
	if e.Login == "" {
		return nil, model.BadRequest("no LOGIN")
	}

	e.FullName = strings.TrimSpace(e.FullName)
	if e.FullName == "" {
		return nil, model.BadRequest("no fullname")
	}

	if e.RoleId != internal.AdminRoleId && e.RoleId != internal.HeadRoleId && e.ClassId == nil {
		return nil, model.BadRequest("NON ADMIN CAN NOT HAVE NO CLASS")

	}

	password, err := GenerateRandomPassword()
	if err != nil {
		return nil, err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := model.User{
		Email:        e.Email,
		Login:        e.Login,
		PasswordHash: passwordHash,
		RoleId:       e.RoleId,
		FullName:     e.FullName,
		Blocked:      false,
		ClassId:      e.ClassId,
		CreatedAt:    time.Now(),
	}

	login, err := s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &model.UserCredentials{Login: login, Password: password}, nil
}

func (s *AdminServiceStruct) UpdateUser(ctx context.Context, e model.CreateAndUpdateUserRequest) error {
	e.Email = strings.TrimSpace(e.Email)
	if e.Email == "" {
		return model.ErrBadRequest
	}

	e.Login = strings.TrimSpace(e.Login)
	if e.Login == "" {
		return model.ErrBadRequest
	}

	e.FullName = strings.TrimSpace(e.FullName)
	if e.FullName == "" {
		return model.ErrBadRequest
	}

	if e.RoleId != internal.AdminRoleId && e.RoleId != internal.HeadRoleId && e.ClassId == nil {
		return model.ErrBadRequest
	}

	if e.Id <= 0 {
		return model.ErrBadRequest
	}

	user := model.User{
		Id:        e.Id,
		Email:     e.Email,
		Login:     e.Login,
		RoleId:    e.RoleId,
		FullName:  e.FullName,
		Blocked:   false,
		ClassId:   e.ClassId,
		CreatedAt: time.Now(),
	}

	err := s.userRepo.UpdateUser(ctx, user)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return model.NotFound(err.Error())
	}
	if err != nil {
		return err
	}

	return nil
}

func (s *AdminServiceStruct) DeleteUser(ctx context.Context, id int) error {
	if id <= 0 {
		return model.ErrBadRequest
	}

	err := s.userRepo.DeleteUser(ctx, id)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return model.ErrNotFound
	}
	if err != nil {
		return err
	}

	return nil
}

func (s *AdminServiceStruct) UpdatePassword(ctx context.Context, userId int) (*model.UserCredentials, error) {
	if userId <= 0 {
		return nil, model.ErrBadRequest
	}

	password, err := GenerateRandomPassword()
	if err != nil {
		return nil, err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	login, err := s.userRepo.UpdateUserPassword(ctx, userId, passwordHash)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, model.NotFound(err.Error())
	}
	if err != nil {
		return nil, err
	}

	return &model.UserCredentials{Login: login, Password: password}, nil
}

func (s *AdminServiceStruct) CreateEquationType(ctx context.Context, e model.CreateEquationTypeRequest) error {
	if e.Class < 1 || e.Class > 4 {
		return model.BadRequest(fmt.Sprintf("class is out of borders from 1 to 4, class: %d", e.Class))
	}

	e.Name = strings.TrimSpace(e.Name)
	if e.Name == "" {
		return model.BadRequest("name is empty")
	}

	e.Description = strings.TrimSpace(e.Description)
	if e.Name == "" {
		return model.BadRequest("description is empty")
	}

	e.Operations = strings.TrimSpace(e.Operations)
	if e.Operations == "" {
		return model.BadRequest("operations is empty ")
	}

	if len(e.Operands) != e.NumOperands {
		return model.BadRequest("numOperands is not the same as len(operands)")
	}

	err := s.equationTypeRepo.CreateFullEquationType(ctx, model.EquationType{
		Class:       e.Class,
		Name:        e.Name,
		Description: e.Description,
		Operations:  e.Operations,
		NumOperands: e.NumOperands,
		NoRemainder: e.NoRemainder,
		MaxResult:   e.MaxResult,
	}, createRequestOperandsToOperands(e.Operands))
	if err != nil {
		return err
	}

	return nil
}

func createRequestOperandsToOperands(op []model.CreateOperandRequest) []model.Operand {
	result := make([]model.Operand, len(op))

	for i, v := range op {
		result[i] = model.Operand{
			MaxValue:     v.MaxValue,
			MinValue:     v.MinValue,
			OperandOrder: v.OperandOrder,
		}
	}

	return result
}

func (s *AdminServiceStruct) DeleteEquationType(ctx context.Context, equationId int) error {
	if equationId <= 0 {
		return model.BadRequest("invalid id")
	}

	return s.equationTypeRepo.DeleteFullEquationType(ctx, equationId)
}

func (s *AdminServiceStruct) UpdateEquationType(ctx context.Context, equationTypeId int, e model.UpdateEquationTypeRequest) error {
	if equationTypeId <= 0 {
		return model.BadRequest("invalid id")
	}

	if e.Class < 1 || e.Class > 4 {
		return model.BadRequest(fmt.Sprintf("class is out of borders from 1 to 4, class: %d", e.Class))
	}

	e.Name = strings.TrimSpace(e.Name)
	if e.Name == "" {
		return model.BadRequest("name is empty")
	}

	e.Description = strings.TrimSpace(e.Description)
	if e.Name == "" {
		return model.BadRequest("description is empty")
	}

	e.Operations = strings.TrimSpace(e.Operations)
	if e.Operations == "" {
		return model.BadRequest("operations is empty ")
	}

	if len(e.Operands) != e.NumOperands {
		return model.BadRequest("numOperands is not the same as len(operands)")
	}

	err := s.equationTypeRepo.UpdateFullEquationType(ctx, model.EquationType{
		Id:          equationTypeId,
		Class:       e.Class,
		Name:        e.Name,
		Description: e.Description,
		Operations:  e.Operations,
		NumOperands: e.NumOperands,
		NoRemainder: e.NoRemainder,
		MaxResult:   e.MaxResult,
	}, updateRequestOperandsToOperands(e.Operands), equationTypeId)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return model.NotFound(fmt.Sprintf("not found row with id: %d, error: %w", equationTypeId, err))
	}
	if err != nil {
		return err
	}

	return nil
}

func updateRequestOperandsToOperands(ops []model.UpdateOperandRequest) []model.Operand {
	result := make([]model.Operand, len(ops))
	for i, v := range ops {
		result[i] = model.Operand{
			Id:           v.Id,
			MaxValue:     v.MaxValue,
			MinValue:     v.MinValue,
			OperandOrder: v.OperandOrder,
		}
	}

	return result
}

func (s *AdminServiceStruct) GetEquationTypes(ctx context.Context) ([]model.EquationType, error) {
	return s.equationTypeRepo.GetEquationTypes(ctx)
}

func (s *AdminServiceStruct) GetOperandsForEquationType(ctx context.Context, equationTypeId int) ([]model.OperandResponse, error) {
	if equationTypeId <= 0 {
		return nil, model.BadRequest("invalid id")
	}

	return s.equationTypeRepo.GetOperandsForEquationTypeId(ctx, equationTypeId)
}

func (s *AdminServiceStruct) JoinEquationTypeAndSection(ctx context.Context, equationTypeId, sectionId int) error {
	if equationTypeId <= 0 || sectionId <= 0 {
		return model.BadRequest("invalid id")
	}

	return s.equationTypeRepo.JoinEquationTypeToSection(ctx, equationTypeId, sectionId)
}

func (s *AdminServiceStruct) UnJoinEquationTypeAndSection(ctx context.Context, equationTypeId, sectionId int) error {
	if equationTypeId <= 0 || sectionId <= 0 {
		return model.BadRequest("invalid id")
	}

	return s.equationTypeRepo.UnJoinEquationTypeToSection(ctx, equationTypeId, sectionId)
}

func (s *AdminServiceStruct) GetSectionsAndEquationTypes(ctx context.Context) ([]model.SectionAndEquationType, error) {
	return s.equationTypeRepo.GetSectionsAndEquationTypes(ctx)
}
