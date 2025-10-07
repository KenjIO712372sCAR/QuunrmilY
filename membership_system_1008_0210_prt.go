// 代码生成时间: 2025-10-08 02:10:28
作者：[您的姓名]
日期：[当前日期]
*/

package main

import (
    "encoding/json"
    "errors"
    "fmt"
    "revel"
    "time"
)

// Member 会员信息结构体
type Member struct {
    ID        int       `json:"id"`        // 会员ID
    Name      string    `json:"name"`     // 会员姓名
    Email     string    `json:"email"`    // 会员邮箱
    CreatedAt time.Time `json:"created_at"` // 创建时间
    UpdatedAt time.Time `json:"updated_at"` // 更新时间
}

// MembershipService 会员服务接口
type MembershipService interface {
    CreateMember(member Member) error
    GetMember(id int) (Member, error)
    UpdateMember(member Member) error
    DeleteMember(id int) error
}

// MemoryMembershipService 内存会员服务实现
type MemoryMembershipService struct {
    members map[int]Member
}

// NewMemoryMembershipService 创建内存会员服务
func NewMemoryMembershipService() *MemoryMembershipService {
    return &MemoryMembershipService{
        members: make(map[int]Member),
    }
}

// CreateMember 创建会员
func (s *MemoryMembershipService) CreateMember(member Member) error {
    if _, exists := s.members[member.ID]; exists {
        return errors.New("会员ID已存在")
    }
    s.members[member.ID] = member
    return nil
}

// GetMember 获取会员信息
func (s *MemoryMembershipService) GetMember(id int) (Member, error) {
    member, exists := s.members[id]
    if !exists {
        return Member{}, errors.New("会员不存在")
    }
    return member, nil
}

// UpdateMember 更新会员信息
func (s *MemoryMembershipService) UpdateMember(member Member) error {
    if _, exists := s.members[member.ID]; !exists {
        return errors.New("会员不存在")
    }
    s.members[member.ID] = member
    return nil
}

// DeleteMember 删除会员
func (s *MemoryMembershipService) DeleteMember(id int) error {
    if _, exists := s.members[id]; !exists {
        return errors.New("会员不存在")
    }
    delete(s.members, id)
    return nil
}

// MembershipController 会员控制器
type MembershipController struct {
    revel.Controller
    service MembershipService
}

// NewMembershipController 创建会员控制器
func NewMembershipController(service MembershipService) *MembershipController {
    return &MembershipController{
        service: service,
    }
}

// AddMember 添加会员
func (c *MembershipController) AddMember() revel.Result {
    var member Member
    if err := json.Unmarshal(c.Params.Form["member"], &member); err != nil {
        return c.RenderJSON(map[string]string{"error": "无效的会员信息"})
    }
    if err := c.service.CreateMember(member); err != nil {
        return c.RenderJSON(map[string]string{"error": err.Error()})
    }
    return c.RenderJSON(map[string]string{"message": "会员添加成功"})
}

// GetMember 获取会员信息
func (c *MembershipController) GetMember(id int) revel.Result {
    member, err := c.service.GetMember(id)
    if err != nil {
        return c.RenderJSON(map[string]string{"error": err.Error()})
    }
    return c.RenderJSON(member)
}

// UpdateMember 更新会员信息
func (c *MembershipController) UpdateMember(id int) revel.Result {
    var member Member
    if err := json.Unmarshal(c.Params.Form["member"], &member); err != nil {
        return c.RenderJSON(map[string]string{"error": "无效的会员信息"})
    }
    member.ID = id
    if err := c.service.UpdateMember(member); err != nil {
        return c.RenderJSON(map[string]string{"error": err.Error()})
    }
    return c.RenderJSON(map[string]string{"message": "会员更新成功"})
}

// DeleteMember 删除会员
func (c *MembershipController) DeleteMember(id int) revel.Result {
    if err := c.service.DeleteMember(id); err != nil {
        return c.RenderJSON(map[string]string{"error": err.Error()})
    }
    return c.RenderJSON(map[string]string{"message": "会员删除成功"})
}

func main() {
    revel.OnAppStart(func() {
        fmt.Println("会员管理系统启动...")
    })

    service := NewMemoryMembershipService()
    controller := NewMembershipController(service)
    revel.InterceptController = controller

    revel.Run()
}
