package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/sheelendar196/go-projects/grpc_project/internal/core/domain"
	"github.com/sheelendar196/go-projects/grpc_project/proto"
	"google.golang.org/grpc"
)

type Service struct {
	proto.UnimplementedEmployeeServiceServer
	Server         *grpc.Server
	empoyeeService EmployeeService
}

func NewService(srv *grpc.Server, empService EmployeeService) *Service {
	return &Service{
		Server:         srv,
		empoyeeService: empService,
	}
}

func (s *Service) Start(ctx context.Context, port string) error {
	listenCfg := net.ListenConfig{}
	l, err := listenCfg.Listen(ctx, "tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return err
	}

	// register health server
	//hs := health.NewServer()
	//grpc_health_v1.RegisterHealthServer(s.Server, hs)
	proto.RegisterEmployeeServiceServer(s.Server, s)
	//reflection.Register(s.Server)
	return s.Server.Serve(l)
}

func (s *Service) Stop() {
	s.Server.GracefulStop()
}

func (s *Service) CreateEmployee(ctx context.Context, req *proto.CreateEmployeeRequest) (*proto.CreateEmployeeResponse, error) {
	err := s.empoyeeService.InsertEmplyee(ctx, getModelEmpployeeByProto(req.GetEmployee()))
	res := &proto.CreateEmployeeResponse{}
	if err != nil {
		res.Success = "false"
		res.Err = "record not create into db"
		return res, err
	}
	res.Success = "true"
	return res, nil
}
func (s *Service) GetEmployee(ctx context.Context, req *proto.GetEmployeeRequest) (*proto.GetEmployeeResponse, error) {
	emp, err := s.empoyeeService.GetEmployeeEntryByID(ctx, req.EmpID)
	res := &proto.GetEmployeeResponse{}
	if err != nil {
		res.Err = "No employee for for this employeeID"
		return res, err
	}
	getProtoResponseByMpdel(res, emp)
	return res, nil
}
func getProtoResponseByMpdel(res *proto.GetEmployeeResponse, emp *domain.Employee) {
	res.Employee = &proto.Employee{
		ID:         emp.ID,
		Name:       emp.Name,
		EmpID:      emp.EmpID,
		Address:    *emp.Address,
		Mobile:     *emp.Mobile,
		Email:      *emp.Email,
		Department: *emp.Department,
	}
}

func getModelEmpployeeByProto(emp *proto.Employee) *domain.Employee {
	employee := &domain.Employee{
		Name:       emp.Name,
		EmpID:      emp.EmpID,
		Address:    &emp.Address,
		Mobile:     &emp.Mobile,
		Email:      &emp.Email,
		Department: &emp.Department,
	}
	return employee
}
