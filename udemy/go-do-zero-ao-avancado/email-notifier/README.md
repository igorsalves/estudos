# Email Notifier - Sistema de Campanhas de Email

## 📋 Visão Geral

Email Notifier é um sistema robusto desenvolvido em Go para criação e gerenciamento de campanhas de email em larga escala. O sistema permite criar, gerenciar e enviar emails para múltiplos destinatários de forma assíncrona, com autenticação via OAuth 2.0 (Keycloak) e processamento em background através de workers.

## 🏗️ Arquitetura

### Visão Arquitetural

O sistema segue os princípios de **Clean Architecture** e **Domain-Driven Design (DDD)**, organizando o código em camadas bem definidas:

```mermaid
graph TB
    subgraph External["🌐 Camada Externa"]
        Client["Cliente HTTP"]
        Worker["Worker<br/>(Background Process)"]
    end
    
    subgraph API["🚀 API REST - Port 3000"]
        Router["Chi Router + Middlewares<br/>(RequestID, Logger, Auth)"]
    end
    
    subgraph Endpoints["📡 Endpoints Layer"]
        Handlers["HTTP Handlers"]
        Auth["Auth Middleware<br/>(OAuth 2.0 / Keycloak)"]
        ErrorHandler["Error Handling"]
    end
    
    subgraph Domain["🎯 Domain Layer (Core Business Logic)"]
        Service["Campaign Service"]
        Entities["Entities<br/>(Campaign, Contact)"]
        Business["Business Rules<br/>(Validations, Status Management)"]
    end
    
    subgraph Infrastructure["🔧 Infrastructure Layer"]
        Repository["Campaign Repository<br/>(GORM)"]
        Database[("PostgreSQL<br/>Database")]
        MailService["Email Service<br/>(SMTP/Gomail)"]
    end
    
    Client -->|HTTP Request| Router
    Router --> Handlers
    Handlers --> Auth
    Auth --> ErrorHandler
    ErrorHandler --> Service
    
    Service --> Entities
    Service --> Business
    Service --> Repository
    Service --> MailService
    
    Repository --> Database
    MailService -->|SMTP| Email["📧 Email Server"]
    
    Worker -->|Polling 60s| Repository
    Worker --> Service
    Worker --> MailService
    
    style Domain fill:#e1f5ff,stroke:#0288d1,stroke-width:3px
    style Infrastructure fill:#fff3e0,stroke:#f57c00,stroke-width:2px
    style Endpoints fill:#f3e5f5,stroke:#7b1fa2,stroke-width:2px
    style API fill:#e8f5e9,stroke:#388e3c,stroke-width:2px
```

### Componentes Principais

#### 1. **API Service** (`cmd/api/main.go`)
- Servidor HTTP REST na porta 3000
- Gerenciamento de rotas com Chi Router
- Middlewares: RequestID, RealIP, Logger, Recoverer
- Autenticação OAuth 2.0 (Keycloak)

#### 2. **Worker Service** (`cmd/worker/main.go`)
- Processo independente para envio assíncrono de emails
- Polling a cada 60 segundos
- Busca campanhas com status "Started"
- Atualiza status após envio (Done/Fail)

#### 3. **Domain Layer** (`internal/domain/campaign`)
- **Entidades**: Campaign, Contact
- **Estados da Campanha**: Pending, Started, Done, Fail, Canceled, Deleted
- **Regras de Negócio**: Validações, transições de estado

#### 4. **Infrastructure Layer**
- **Database**: PostgreSQL com GORM
- **Email**: SMTP via Gomail
- **Auth**: Keycloak (OpenID Connect)

## 🔧 Tecnologias Utilizadas

| Tecnologia | Versão | Propósito |
|------------|--------|-----------|
| Go | 1.25.5 | Linguagem principal |
| Chi Router | v5.2.3 | Roteamento HTTP |
| GORM | v1.31.1 | ORM para PostgreSQL |
| PostgreSQL | 15 | Banco de dados |
| Keycloak | 21.1.1 | Autenticação OAuth 2.0 |
| Gomail | v2 | Envio de emails SMTP |
| Docker Compose | 3.8 | Orquestração de containers |
| godotenv | v1.5.1 | Gerenciamento de variáveis de ambiente |

## 📁 Estrutura de Diretórios

```
email-notifier/
├── cmd/
│   ├── api/                    # Entry point da API REST
│   │   └── main.go
│   └── worker/                 # Entry point do Worker
│       └── main.go
├── internal/
│   ├── contract/               # DTOs e contratos
│   │   ├── CampaignResponse.go
│   │   └── NewCampaignDto.go
│   ├── domain/                 # Camada de domínio
│   │   └── campaign/
│   │       ├── campaign.go     # Entidade Campaign
│   │       ├── service.go      # Lógica de negócio
│   │       └── repository.go   # Interface de repositório
│   ├── endpoints/              # Handlers HTTP
│   │   ├── auth.go             # Middleware de autenticação
│   │   ├── campaigns_*.go      # Handlers de campanhas
│   │   └── handler_error.go    # Tratamento de erros
│   ├── infrastructure/         # Infraestrutura
│   │   ├── database/           # Implementação do repositório
│   │   └── mail/               # Serviço de email
│   ├── internal-errors/        # Gestão de erros
│   └── test/                   # Mocks e testes
├── docker-compose.yml          # Orquestração Docker
├── go.mod                      # Dependências
└── README.md                   # Documentação
```

## 🔐 Segurança e Autenticação

### Fluxo de Autenticação

1. Cliente obtém token JWT do Keycloak
2. Token é enviado no header `Authorization: Bearer {token}`
3. Middleware `Auth` valida o token via OIDC
4. Email do usuário é extraído do token e injetado no contexto
5. Requisição é processada com identificação do usuário

### Configuração do Keycloak

- **URL**: http://localhost:8080
- **Realm**: Configurável
- **Client ID**: `emailn`
- **Admin**: admin/admin (desenvolvimento)

## 📡 API Endpoints

### Base URL
```
http://localhost:3000
```

### Endpoints Disponíveis

#### 1. Health Check
```http
GET /ping
```
**Resposta**: `pong`

---

#### 2. Criar Campanha
```http
POST /campaigns
Authorization: Bearer {token}
Content-Type: application/json
```

**Body**:
```json
{
  "name": "Black Friday 2026",
  "content": "Aproveite as ofertas exclusivas!",
  "emails": [
    "cliente1@example.com",
    "cliente2@example.com"
  ]
}
```

**Resposta** (201):
```json
{
  "id": "campaign_id_generated"
}
```

**Validações**:
- `name`: mínimo 5, máximo 24 caracteres
- `content`: mínimo 5, máximo 1024 caracteres
- `emails`: mínimo 1 email válido

---

#### 3. Buscar Campanha por ID
```http
GET /campaigns/{id}
Authorization: Bearer {token}
```

**Resposta** (200):
```json
{
  "id": "c9s8d9f8s",
  "name": "Black Friday 2026",
  "content": "Aproveite as ofertas exclusivas!",
  "status": "Pending",
  "amountOfEmailsToSend": 2,
  "createdBy": "user@example.com"
}
```

---

#### 4. Iniciar Campanha
```http
PATCH /campaigns/start/{id}
Authorization: Bearer {token}
```

**Comportamento**:
- Valida se a campanha está com status "Pending"
- Altera status para "Started"
- Processa envio de forma assíncrona via goroutine

**Resposta** (200):
```json
{
  "message": "Campaign started successfully"
}
```

---

#### 5. Deletar Campanha
```http
DELETE /campaigns/{id}
Authorization: Bearer {token}
```

**Comportamento**:
- Valida se a campanha está com status "Pending"
- Marca campanha e contatos como deletados (soft delete)

**Resposta** (200):
```json
{
  "message": "Campaign deleted successfully"
}
```

## 🎯 Domain Model

### Entidade Campaign

```go
type Campaign struct {
    ID        string      // UUID gerado com xid
    Name      string      // Nome da campanha (5-24 chars)
    CreatedOn time.Time   // Data de criação
    UpdatedOn time.Time   // Data de atualização (auto)
    Content   string      // Conteúdo do email (5-1024 chars)
    Contacts  []Contact   // Lista de contatos (min 1)
    Status    string      // Estado atual
    CreatedBy string      // Email do criador
}

type Contact struct {
    ID         string  // UUID gerado com xid
    Email      string  // Email do contato
    CampaignId string  // Chave estrangeira
}
```

### Máquina de Estados

```mermaid
stateDiagram-v2
    [*] --> Pending: Nova Campanha
    
    Pending --> Started: Start()
    Pending --> Deleted: Delete()
    
    Started --> Done: Envio com Sucesso
    Started --> Fail: Erro no Envio
    
    Done --> [*]
    Fail --> [*]
    Deleted --> [*]
    
    note right of Pending
        Estado inicial
        Aceita: Start, Delete
    end note
    
    note right of Started
        Processamento assíncrono
        Worker envia emails
    end note
    
    note right of Done
        Emails enviados
        Estado final
    end note
    
    note right of Fail
        Erro no envio
        Estado final
    end note
```

**Regras de Transição**:
- Apenas campanhas "Pending" podem ser iniciadas ou deletadas
- "Started" → Worker processa → "Done" ou "Fail"
- Transições são irreversíveis

## 🔄 Fluxo de Processamento

### Criação de Campanha

```
1. Cliente → POST /campaigns (com token JWT)
2. Auth Middleware valida token e extrai email
3. Handler recebe dados e chama CampaignService
4. Service valida dados e cria entidade Campaign
5. Repository persiste no PostgreSQL
6. Retorna ID da campanha ao cliente
```

### Envio de Emails (Assíncrono)

```mermaid
sequenceDiagram
    participant Client
    participant API
    participant Database
    participant Worker
    participant SMTP

    Client->>API: PATCH /campaigns/start/{id}
    API->>Database: GetBy(id)
    Database-->>API: Campaign (status: Pending)
    API->>API: Validate Status
    API->>Database: Update Status = "Started"
    API-->>Client: 200 OK

    Note over Worker: Loop a cada 60 segundos
    
    loop Polling
        Worker->>Database: GetCampaignsToBeSent()
        Database-->>Worker: Campaigns (status: Started, updated_on > 1 min)
        
        loop For Each Campaign
            Worker->>Worker: Process Campaign
            Worker->>SMTP: SendMail(campaign)
            
            alt Email Enviado com Sucesso
                SMTP-->>Worker: Success
                Worker->>Database: Update Status = "Done"
            else Erro no Envio
                SMTP-->>Worker: Error
                Worker->>Database: Update Status = "Fail"
            end
        end
        
        Worker->>Worker: Sleep 60s
    end
```

**Critério de Seleção do Worker**:
```sql
WHERE status = 'Started' 
AND date_part('minute', now() - updated_on) >= 1
```

## 🐳 Docker e Infraestrutura

### Variáveis de Ambiente Necessárias (.env)

```bash
# Database
DATABASE_URL=postgres://postgres:postgres@localhost:5432/email_notifier

# Keycloak
KEYCLOACK=http://localhost:8080/realms/{realm_name}

# SMTP
EMAIL_SMTP=smtp.gmail.com
EMAIL_USER=your-email@gmail.com
EMAIL_PASSWORD=your-app-password
```

## 🚀 Como Executar

### Pré-requisitos
- Go 1.25.5+
- Docker e Docker Compose
- Air (para hot reload) - opcional

### Passo a Passo

1. **Clone o repositório e navegue até o diretório**
```bash
cd email-notifier
```

2. **Configure as variáveis de ambiente**
```bash
cp .env.example .env
# Edite o arquivo .env com suas credenciais
```

3. **Inicie os serviços de infraestrutura**
```bash
docker-compose up -d
```

4. **Configure o Keycloak**
- Acesse http://localhost:8080
- Crie um realm
- Crie um client com ID `emailn`
- Configure usuários

5. **Execute as migrations (se necessário)**
```bash
# GORM criará as tabelas automaticamente
```

6. **Inicie a API**
```bash
# Com Air (hot reload)
cd cmd/api
air

# Ou manualmente
go run cmd/api/main.go
```

7. **Inicie o Worker** (em outro terminal)
```bash
go run cmd/worker/main.go
```

## 🧪 Testes

O projeto inclui testes unitários para:
- Domain Layer (campaign_test.go, service_test.go)
- Endpoints (campaigns_*_test.go)
- Error Handling (handler_error_test.go)

### Executar Testes
```bash
# Todos os testes
go test ./...

# Testes com cobertura
go test -cover ./...

# Testes de um pacote específico
go test ./internal/domain/campaign/...
```

### Mocks
Localizados em `internal/test/internal-mock/`:
- `campaign_repository_mock.go`
- `campaign_service_mock.go`

## 📊 Padrões e Boas Práticas

### Clean Architecture
- **Separação de Responsabilidades**: Cada camada tem uma responsabilidade clara
- **Dependency Inversion**: Camadas internas não dependem de camadas externas
- **Testabilidade**: Interfaces permitem mock de dependências

### Design Patterns Utilizados

#### 1. Repository Pattern
```go
type Repository interface {
    Create(campaign *Campaign) error
    GetBy(id string) (*Campaign, error)
    Update(campaign *Campaign) error
    Delete(campaign *Campaign) error
}
```

#### 2. Service Layer Pattern
```go
type Service interface {
    Create(newCampaign contract.NewCampaign) (string, error)
    GetBy(id string) (*contract.CampaignResponse, error)
    Delete(id string) error
    Start(id string) error
}
```

#### 3. DTO (Data Transfer Object)
- `NewCampaignDto.go`: Entrada de dados
- `CampaignResponse.go`: Saída de dados

#### 4. Error Handling Middleware
```go
func HandlerError(handler func(w http.ResponseWriter, r *http.Request) error) 
    http.HandlerFunc
```

## 🔒 Tratamento de Erros

### Erros de Domínio
```go
var ErrInternal error = errors.New("Internal Server Error")
```

### Erros HTTP
- 400: Bad Request (validações)
- 401: Unauthorized (token inválido)
- 404: Not Found (campanha não encontrada)
- 500: Internal Server Error

## 📝 Licença

Projeto desenvolvido para fins educacionais.

---

**Desenvolvido com ❤️ em Go**
