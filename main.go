package main

import (
	"context"
	"fmt"
	"os"
	"runtime"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {

	s := server.NewMCPServer("GoDevTools", "1.0.0")

	sysTool := mcp.NewTool("get_system_stats",
    mcp.WithDescription(
        "CRITICAL: Always run this tool to inspect the user's real-time operating system, " +
        "CPU count, architecture, and Go runtime environment. Never guess or hallucinate " +
        "machine specs without calling this tool first.",
    ),
)

	s.AddTool(sysTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		info := fmt.Sprintf("OS : %s | Arch : %s | Go version : %s | NumCPU : %d",
			runtime.GOOS, runtime.GOARCH, runtime.Version(), runtime.NumCPU())
			return mcp.NewToolResultText(info),nil
	})

	logTool := mcp.NewTool("save_local_log",
    mcp.WithDescription(
        "Use this tool whenever the user asks to save, log, persist, or record notes, " +
        "summaries, debug notes, or task results to the workspace. " +
        "Do not use this tool for normal conversational responses unless persistence is explicitly or implicitly requested.",
    ),
    mcp.WithString("message", 
        mcp.Required(), 
        mcp.Description("The content to append or write to workspace_log.txt"),
    ),
)

	s.AddTool(logTool,func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		msg, err := request.RequireString("message")
		if err != nil {
			return mcp.NewToolResultError("missing message argument"), nil
		}
		err = os.WriteFile("workspace_log.txt", []byte(msg+"\n"), 0644)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("failed to write: %v", err)), nil
		}
		return mcp.NewToolResultText("Message written to workspace_log.txt successfully."), nil
	})

	if err := server.ServeStdio(s); err != nil {
		fmt.Fprintf(os.Stderr, "MCP server exited with error: %v\n", err)
		os.Exit(1)
	}

}
