# golang-ai-agent

A Go implementation of AI agent patterns based on Anthropic's guide for building effective agents.

## About

This repository implements the principles and patterns outlined in Anthropic's excellent guide: [Building Effective Agents](https://www.anthropic.com/engineering/building-effective-agents). The goal is to create a practical, iterative implementation in Go that demonstrates the core concepts of effective AI agent development.

The project follows Anthropic's key recommendations:
- Start simple and iterate
- Focus on transparency and clear interfaces
- Build with extensibility in mind
- Implement proper tool design patterns

## Current State

This is currently a basic chatbot implementation with:
- OpenAI API integration
- Conversation memory
- Environment-based configuration
- Simple CLI interface

## Features

- **OpenAI Integration**: Uses the official go-openai library for LLM communication
- **Conversation Memory**: Maintains chat history for contextual conversations
- **Environment Configuration**: Secure API key management via .env files
- **Interactive CLI**: Simple command-line interface with graceful exit

## Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/luc-pimentel/golang-ai-agent.git
   cd golang-ai-agent
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Create a `.env` file with your OpenAI API key:
   ```
   OPENAI_API_KEY=your_openai_api_key_here
   ```

4. Run the chatbot:
   ```bash
   go run main.go
   ```

## Usage

Once running, simply type messages to chat with the AI. Type `quit` to exit.

```
AI Chatbot initialized. Type 'quit' to exit.
> Hello, how are you?
Bot: I'm doing well, thank you! How can I help you today?
> quit
Goodbye!
```

## Roadmap

This implementation will be expanded iteratively to include:
- Tool calling capabilities
- Workflow patterns (chaining, routing, orchestration)
- Retrieval augmented generation (RAG)
- Agent memory systems
- Multi-agent orchestration

## Credits

- Based on [Building Effective Agents](https://www.anthropic.com/engineering/building-effective-agents) by Anthropic
- Uses [go-openai](https://github.com/sashabaranov/go-openai) for OpenAI API integration
- Environment management via [godotenv](https://github.com/joho/godotenv)

## License

MIT License - see LICENSE file for details.