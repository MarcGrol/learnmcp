package handlers

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"github.com/MarcGrol/service-catalog-mcp-server/internal/resp"
	
)

// NewListModulesByComplexityTool returns the MCP tool definition and its handler for listing modules.
func (h *MCPHandler) NewListModulesByComplexityTool() server.ServerTool {
	return server.ServerTool{
		Tool: mcp.NewTool(
			"list_modules_by_complexity",
			mcp.WithDescription("Lists all modules in the catalog ordered DESC on complexity limited up to limit_to modules."),
			mcp.WithNumber("limit_to", mcp.Description("Maximum number of modules to return.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			// extract params
			limit := request.GetInt("limit_to", 20)

			// call business logic
			modules, err := h.repo.ListModulesByCompexity(ctx, limit)
			if err != nil {
				return mcp.NewToolResultError(
					resp.InternalError(ctx,
						fmt.Sprintf("error listing modules by complexity: %s", err))), nil
			}

			results := []moduleDescriptor{}
			for _, mod := range modules {
				results = append(results, moduleDescriptor{
					ModuleID:        mod.ModuleID,
					Name:            mod.Name,
					Description:     mod.Description,
					ComplexityScore: mod.ComplexityScore,
				})
			}
			return mcp.NewToolResultText(resp.Success(ctx, results)), nil
		},
	}
}
