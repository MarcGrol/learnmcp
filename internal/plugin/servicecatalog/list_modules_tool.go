package servicecatalog

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"github.com/MarcGrol/service-catalog-mcp-server/internal/core/resp"
)

// NewListModulesTool returns the MCP tool definition and its handler for listing modules.
func (h *mcpHandler) listModulesTool() server.ServerTool {
	return server.ServerTool{
		Tool: mcp.NewTool(
			"list_modules",
			mcp.WithDescription("Lists all modules in the catalog."),
			mcp.WithString("filter_keyword", mcp.Required(), mcp.Description("The keyword to filter modules by.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			// extract params
			keyword, err := request.RequireString("filter_keyword")
			if err != nil {
				return mcp.NewToolResultError(
					resp.InvalidInput(ctx, "Missing filter_keyword",
						"filter_keyword",
						"Use a non-empty string as keyword")), nil
			}

			// call business logic
			modules, err := h.repo.ListModules(ctx, keyword)
			if err != nil {
				return mcp.NewToolResultError(
					resp.InternalError(ctx,
						fmt.Sprintf("error listing modules with keyword %s: %s", keyword, err))), nil
			}

			results := []moduleDescriptor{}
			for _, mod := range modules {
				results = append(results, moduleDescriptor{
					ModuleID:    mod.ModuleID,
					Name:        mod.Name,
					Description: mod.Description,
				})
			}
			return mcp.NewToolResultText(resp.Success(ctx, results)), nil
		},
	}
}

type moduleDescriptor struct {
	ModuleID        string
	Name            string
	Description     string
	ComplexityScore float32 `json:",omitempty"`
}
