package interfacego

import "fmt"

// 实际应用场景1：数据存储抽象
type UserRepository interface {
	Save(user User) error
	FindByID(id int) (User, error)
	Delete(id int) error
}

type User struct {
	ID   int
	Name string
}

// 不同的实现
type MySQLUserRepo struct {
	connectionString string
}

func (m MySQLUserRepo) Save(user User) error {
	fmt.Printf("保存用户 %s 到MySQL数据库\n", user.Name)
	return nil
}

func (m MySQLUserRepo) FindByID(id int) (User, error) {
	fmt.Printf("从MySQL查找用户ID: %d\n", id)
	return User{ID: id, Name: "MySQL用户"}, nil
}

func (m MySQLUserRepo) Delete(id int) error {
	fmt.Printf("从MySQL删除用户ID: %d\n", id)
	return nil
}

type MongoUserRepo struct {
	database string
}

func (m MongoUserRepo) Save(user User) error {
	fmt.Printf("保存用户 %s 到MongoDB\n", user.Name)
	return nil
}

func (m MongoUserRepo) FindByID(id int) (User, error) {
	fmt.Printf("从MongoDB查找用户ID: %d\n", id)
	return User{ID: id, Name: "MongoDB用户"}, nil
}

func (m MongoUserRepo) Delete(id int) error {
	fmt.Printf("从MongoDB删除用户ID: %d\n", id)
	return nil
}

// 业务逻辑层不关心具体实现
type UserService struct {
	repo UserRepository // 依赖接口，不是具体实现
}

func (s UserService) CreateUser(name string) {
	user := User{ID: 1, Name: name}
	s.repo.Save(user)
	fmt.Printf("用户 %s 创建成功\n", name)
}

func (s UserService) GetUser(id int) {
	user, _ := s.repo.FindByID(id)
	fmt.Printf("获取到用户: %s\n", user.Name)
}

func DependencyInversionDemo() {
	fmt.Println("=== 依赖倒置原则演示 ===")

	// 可以随时切换不同的实现
	fmt.Println("1. 使用MySQL实现:")
	mysqlRepo := MySQLUserRepo{connectionString: "mysql://..."}
	userService1 := UserService{repo: mysqlRepo}
	userService1.CreateUser("张三")
	userService1.GetUser(1)

	fmt.Println("\n2. 切换到MongoDB实现:")
	mongoRepo := MongoUserRepo{database: "userdb"}
	userService2 := UserService{repo: mongoRepo}
	userService2.CreateUser("李四")
	userService2.GetUser(1)

	fmt.Println("\n优势：业务逻辑完全不变，只是换了存储方式！")
}
