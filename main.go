package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

type Pet struct {
	Name   string
	Sex    string
	Intact bool
	Age    string
	Breed  string
}

type Node struct {
	Name      string
	Extension string
	Folders   []*Node
	File      []*Node

}

type BasicConfig struct {
	ProjectName   string
	JWT           bool
	Swagger       bool
	Redis         bool
	Validator     bool
	DBPostgres    bool
	DBMysql       bool
	DBSqlServer   bool
	EchoFrameWork bool
	GinFrameWork  bool
}

var templateRegistry = map[string]string{
	".env":                "env.tmpl",
	"env.go":              "util_env.tmpl",
	"stringutils.go":      "stringutils.tmpl",
	"config_models.go":    "config_models.tmpl",
	"config.go":           "config.tmpl",
	"main.go":             "main.tmpl",
	"database.go":         "database.tmpl",
	"factory.go":          "factory.tmpl",
	"redis.go":            "redis.tmpl",
	"middleware.go":       "middleware.tmpl",
	"auth.go":             "auth.tmpl",
	"response_model.go":   "response_model.tmpl",
	"error_handler.go":    "error_handler.tmpl",
	"success_handler.go":  "success_handler.tmpl",
	"pagination.go":       "pagination.tmpl",
	"server_routes.go":    "server_routes.tmpl",
	".gitignore":          "gitignore.tmpl",
	"README.md":           "readme.tmpl",
	"controller.go":       "controller.tmpl",
	"service.go":          "service.tmpl",
	"repository.go":       "repository.tmpl",
	"model.go":            "model.tmpl",
	"dto.go":              "dto.tmpl",
	"routes.go":           "routes.tmpl",
	"entity.go":           "entity.tmpl",
	"validator.go":        "validator.tmpl",
	"custom_validator.go": "custom_validator.tmpl",
	"jwt.go": "token.tmpl",
}

func (n *Node) GetFileWithExtension() string {
	return fmt.Sprintf("%s.%s", n.Name, n.Extension)
}

func (n *Node) CreateFile(root string, config *BasicConfig) {
	tmplFilename, ok := templateRegistry[n.GetFileWithExtension()]
	if !ok {
		fmt.Println("--->", n.GetFileWithExtension())
		return
	}

	if strings.Contains(tmplFilename, "redis") && !config.Redis {
		return
	}
	if strings.Contains(tmplFilename, "swagger") && !config.Swagger {
		return
	}
	if strings.Contains(tmplFilename, "auth") && !config.JWT {
		return
	}
	if strings.Contains(tmplFilename, "validator") && !config.Validator {
		return
	}

	tmpl := template.Must(template.New(tmplFilename).ParseFiles(tmplFilename))
	filename := filepath.Join(root, n.GetFileWithExtension())
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(config.ProjectName)
	err = tmpl.Execute(file, config)
	if err != nil {
		panic(err)
	}
}

func (n *Node) CreateNode(root string, config *BasicConfig) {
	fmt.Println(root)
	if n.Extension == "" {
		currentPath := filepath.Join(root, n.Name)
		os.MkdirAll(currentPath, os.ModePerm)
		if len(n.File) > 0 {
			for _, child := range n.File {
				child.CreateFile(currentPath, config)
			}
		}
		if len(n.Folders) > 0 {
			for _, child := range n.Folders {
				child.CreateNode(currentPath, config)
			}
		}
		return
	}

	// create file
	n.CreateFile(root, config)
}
func testCreateFolder() {
	root := "generated"
	// Remove the directory if it exists
	if err := os.RemoveAll(root); err != nil {
		fmt.Println("Error removing directory:", err)
		return
	}
	projectName := "demo"
	os.MkdirAll(root, os.ModePerm)
	node := Node{
		Name: projectName,
		Folders: []*Node{
			{
				Name: "cmd",
				Folders: []*Node{
					{
						Name: "cron",
					},
				},
			},
			{
				Name: "migrations",
			},
			{
				Name: "test",
			},
			{
				Name: "internals",
				Folders: []*Node{
					{
						Name: "factory",
						File: []*Node{
							{
								Name:      "factory",
								Extension: "go",
							},
						},
					},
					{
						Name: "abstraction",
						File: []*Node{
							{
								Name:      "pagination",
								Extension: "go",
							},
							{
								Name:      "entity",
								Extension: "go",
							},
						},
					},
					{
						Name: "utils",
						Folders: []*Node{
							{
								Name: "env",
								File: []*Node{
									{
										Name:      "env",
										Extension: "go",
									},
								},
							},
							{
								Name: "response",
								File: []*Node{
									{
										Name:      "response_model",
										Extension: "go",
									},
									{
										Name:      "success_handler",
										Extension: "go",
									},
									{
										Name:      "error_handler",
										Extension: "go",
									},
								},
							},
							{
								Name: "validator",
								File: []*Node{
									{
										Name:      "validator",
										Extension: "go",
									},
									{
										Name:      "custom_validator",
										Extension: "go",
									},
								},
							},
							{
								Name: "token",
								File: []*Node{
									{
										Name: "jwt",
										Extension: "go",
									},
								},
							},
						},
					},
					{
						Name: "middleware",
						File: []*Node{
							{
								Name:      "middleware",
								Extension: "go",
							},
							{
								Name:      "auth",
								Extension: "go",
							},
						},
					},
					{
						Name: "server",
						File: []*Node{
							{
								Name:      "server_routes",
								Extension: "go",
							},
						},
					},
					{
						Name: "pkg",
						Folders: []*Node{
							{
								Name: "database",
								File: []*Node{
									{
										Name:      "database",
										Extension: "go",
									},
								},
							},
							{
								Name: "redisutil",
								File: []*Node{
									{
										Name:      "redis",
										Extension: "go",
									},
								},
							},
						},
					},
					{
						Name: "app",
						Folders: []*Node{
							{
								Name: "example_feat",
								File: []*Node{
									{
										Name:      "controller",
										Extension: "go",
									},
									{
										Name:      "service",
										Extension: "go",
									},
									{
										Name:      "repository",
										Extension: "go",
									},
									{
										Name:      "model",
										Extension: "go",
									},
									{
										Name:      "dto",
										Extension: "go",
									},
									{
										Name:      "routes",
										Extension: "go",
									},
								},
							},
						},
					},
				},
			},
			{
				Name: "config",
				File: []*Node{
					{
						Name:      "config",
						Extension: "go",
					},
					{
						Name:      "stringutils",
						Extension: "go",
					},
					{
						Name:      "config_models",
						Extension: "go",
					},
				},
			},
		},
		File: []*Node{
			{
				Name:      "",
				Extension: "env",
			},
			{
				Name:      "",
				Extension: "gitignore",
			},
			{
				Name:      "README",
				Extension: "md",
			},
			{
				Name:      "main",
				Extension: "go",
			},
		},
	}
	projectDir := root + `\` + projectName
	config := &BasicConfig{
		ProjectName: "demo",
		DBPostgres:  true,
		JWT:         true,
		Swagger:     true,
		// Redis:         true,
		Validator:     true,
		EchoFrameWork: true,
	}
	node.CreateNode(root, config)
	if config.Swagger {
		runSwagInit(projectDir, `.\\docs`)
	}
	initProject(projectDir, projectName)
}

func runSwagInit(projectDir string, outputDir string) {
	fmt.Println("Generating swagger docs in", projectDir)
	if err := runCommand(projectDir, "swag", "init", "-o", outputDir); err != nil {
		fmt.Println("Error running swag init:", err)
		return
	}
}
func testCreateNode() {
	root := "generated"
	os.MkdirAll(root, os.ModePerm)
	node := Node{
		Name:      "",
		Extension: "env",
		// Detail:    &FileDetail{Purpose: PURPOSE_ENV},
	}
	node.CreateNode(root, &BasicConfig{})
}

func testTemplate() {
	baseDir := "output"
	os.MkdirAll(baseDir, os.ModePerm)

	dogs := []Pet{
		{
			Name:   "Jujube",
			Sex:    "Female",
			Intact: false,
			Age:    "10 months",
			Breed:  "German Shepherd/Pitbull",
		},
		{
			Name:   "Zephyr",
			Sex:    "Male",
			Intact: true,
			Age:    "13 years, 3 months",
			Breed:  "German Shepherd/Border Collie",
		},
	}
	// var tmplFile = "pets.tmpl"
	var tmplFile = "go.tmpl"
	tmpl, err := template.New(tmplFile).ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}

	filename := filepath.Join(baseDir, "main.go")
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = tmpl.Execute(file, dogs)
	if err != nil {
		panic(err)
	}
}

func runCommand(dir, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir // Set working directory
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func initProject(projectDir string, moduleName string) {
	// Ensure the directory exists
	if _, err := os.Stat(projectDir); os.IsNotExist(err) {
		fmt.Println("Project directory does not exist:", projectDir)
		return
	}

	// Run `go mod init <module-name>` in the specified directory
	fmt.Println("Initializing Go module in", projectDir)
	if err := runCommand(projectDir, "go", "mod", "init", moduleName); err != nil {
		fmt.Println("Error initializing module:", err)
		return
	}

	// Run `go mod tidy` in the specified directory
	fmt.Println("Running go mod tidy in", projectDir)
	if err := runCommand(projectDir, "go", "mod", "tidy"); err != nil {
		fmt.Println("Error running go mod tidy:", err)
		return
	}

	// format golang code
	fmt.Println("Running go fmt in", projectDir)
	if err := runCommand(projectDir, "go", "fmt", "./..."); err != nil {
		fmt.Println("Error running go format:", err)
		return
	}

	fmt.Println("Go module setup completed successfully in", projectDir)
}
func main() {
	// testCreateNode()
	testCreateFolder()
} // end main
