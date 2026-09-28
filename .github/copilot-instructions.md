# Tool Calling Rules
- Whenever the user asks any question about their computer, operating system, hardware, CPU, or development environment, you MUST invoke the `get_system_stats` tool from `my-go-server` before answering.
- Do not make assumptions or answer from general knowledge when a local MCP tool is available.

# Logging and Note-Taking Policy
- When the user asks to "wrap up", "summarize session", or "record this decision", invoke `save_local_log` with a structured bulleted summary before completing your response.
- Do not call `save_local_log` for regular code explanations or trivial questions.