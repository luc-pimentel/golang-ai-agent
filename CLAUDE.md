# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go implementation of AI agent patterns based on Anthropic's guide for building effective agents. The project follows an iterative development approach, starting with a basic chatbot and expanding to include advanced agent capabilities.

## Core Architecture

**Current Implementation (Basic Chatbot)**:
- Single `main.go` file with `Chatbot` struct
- OpenAI API integration via `go-openai` library  
- In-memory conversation history stored as slice of `ChatCompletionMessage`
- Environment-based configuration using `.env` files

**Design Philosophy**:
- Start simple and iterate (following Anthropic's recommendations)
- Maintain conversation memory within session
- Use standard Go patterns and minimal dependencies
- Clear separation between API client and conversation logic

## Essential Commands

**Run the application**:
```bash
go run main.go
```

**Install/update dependencies**:
```bash
go mod download
go mod tidy
```

**Build binary**:
```bash
go build -o chatbot main.go
```

## Configuration Requirements

- Requires `.env` file with `OPENAI_API_KEY` variable
- Uses GPT-3.5-turbo model by default
- Conversation history persists only during runtime (no external storage)

## Future Architecture (Roadmap)

The codebase is designed to evolve toward:
- Tool calling capabilities
- Workflow patterns (chaining, routing, orchestration) 
- Retrieval augmented generation (RAG)
- Agent memory systems
- Multi-agent orchestration

When implementing new features, maintain the simple, iterative approach and follow Anthropic's agent design principles from their Building Effective Agents guide.