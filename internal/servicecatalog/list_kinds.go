package servicecatalog

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"github.com/MarcGrol/service-catalog-mcp-server/internal/resp"
	
)

// NewListKindsTool returns the MCP tool definition and its handler for listing kinds.
func (h *mcpHandler) NewListKindsTool() server.ServerTool {
	return server.ServerTool{
		Tool: mcp.NewTool(
			"list_kinds",
			mcp.WithDescription("Lists all module kinds in the catalog."),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

			// call business logic
			flows, err := h.repo.ListKinds(ctx)
			if err != nil {
				return mcp.NewToolResultError(
					resp.InternalError(ctx,
						fmt.Sprintf("error listing kinds: %s", err))), nil
			}

			return mcp.NewToolResultText(resp.Success(ctx, flows)), nil
		},
	}
}
