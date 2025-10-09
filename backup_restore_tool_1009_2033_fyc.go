// 代码生成时间: 2025-10-09 20:33:56
package main

import (
    "encoding/json"
    "log"
    "net/http"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
    "time"
    "revel"
)

// BackupSchedule is a structure to store backup schedule details.
type BackupSchedule struct {
    Frequency string `json:"frequency"`
    Time      string `json:"time"`
}

// BackupResponse is a structure to store the backup response details.
type BackupResponse struct {
    Status  string `json:"status"`
    Message string `json:"message"`
}

func main() {
    // Initialize Revel framework.
    revel.OnAppStart(InitDB)
    revel.FilterChain.Clear()
    revel.FilterChain.Add(revel.PanicFilter)
    revel.FilterChain.Add(revel.ActionInvoker)
    revel.StartServer()
}

// InitDB is called when the application starts.
func InitDB() {
    // Initialize database connection here if needed.
}

// BackupHandler handles the system backup.
func (c *App) BackupHandler() revel.Result {
    // Define the backup directory path.
    backupPath := "./backups"
    err := os.MkdirAll(backupPath, os.ModePerm)
    if err != nil {
        return c.RenderError(err)
    }

    // Define the backup file name with timestamp.
    timestamp := time.Now().Format("20060102150405")
    backupFileName := filepath.Join(backupPath, "system_backup_" + timestamp + ".tar.gz")

    // Create the backup command.
    cmd := exec.Command("tar", "-czf", backupFileName, "./")

    // Run the backup command and handle errors.
    output, err := cmd.CombinedOutput()
    if err != nil {
        log.Printf("Backup error: %s
Output: %s
", err, string(output))
        return c.RenderJSON(BackupResponse{Status: "error", Message: "Failed to create backup."})
    }

    // Return a success response.
    return c.RenderJSON(BackupResponse{Status: "success", Message: "Backup created successfully."})
}

// RestoreHandler handles the system restore.
func (c *App) RestoreHandler() revel.Result {
    // Define the restore directory path and file name.
    backupPath := "./backups"
    backupFileName := "system_backup.tar.gz" // You should modify this to accept the specific backup file.

    // Check if the backup file exists.
    _, err := os.Stat(filepath.Join(backupPath, backupFileName))
    if os.IsNotExist(err) {
        return c.RenderJSON(BackupResponse{Status: "error", Message: "Backup file not found."})
    }

    // Create the restore command.
    cmd := exec.Command("tar", "-xzf", filepath.Join(backupPath, backupFileName), "-C", "./")

    // Run the restore command and handle errors.
    output, err := cmd.CombinedOutput()
    if err != nil {
        log.Printf("Restore error: %s
Output: %s
", err, string(output))
        return c.RenderJSON(BackupResponse{Status: "error", Message: "Failed to restore system."})
    }

    // Return a success response.
    return c.RenderJSON(BackupResponse{Status: "success", Message: "System restored successfully."})
}

// RenderError is a helper function to render error responses.
func (c *App) RenderError(err error) revel.Result {
    return c.RenderJSON(BackupResponse{Status: "error", Message: err.Error()})
}
