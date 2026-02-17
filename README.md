## Abakcus: Small Business Tools

### Development

The backend lives in the `backend` directory and is a Go module.
To run the API locally you need a MongoDB instance and the following
environment variables (see `backend/.env` for an example):

```
MONGO_URI=mongodb://localhost:27017
MONGO_DB=abakcus
```

Start the service with:s

```sh
cd backend
go run ./cmd/api
```

Connection attempts are logged on startup and the process will exit if the
MongoDB handshake fails.

### Installation

#### 1. Smart Contracts

Navigate to the `contract` folder to install dependencies and deploy contracts.

```bash
cd contract
npm install
```

Compile the contracts:
```bash
npx hardhat compile
```

Deploy to the network (e.g., Primordial):
```bash
npx hardhat run scripts/deploy.js --network primordial
# OR using the package.json script
npm run deploy
```
*Note the deployed `TokenFactory` address and update `FACTORY_CONTRACT_ADDRESS` in your frontend `.env`.*

#### 2. Frontend Application

Navigate to the `frontend` folder.

```bash
cd frontend
npm install
```

Run the development server:

```bash
npm run dev
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

## 📚 Documentation

For detailed documentation on how the code works, please refer to [DOCS.md](./DOCS.md).

## 🎨 Styling Guidelines

We follow specific styling guidelines to maintain consistency. See [STYLING_GUIDELINE.md](./STYLING_GUIDELINE.md).

## 🤝 Contributing

We welcome contributions! Please check [CONTRIBUTING.md](./CONTRIBUTING.md) for details on how to contribute to this project.

## 🐛 Issues

Found a bug or have a suggestion? Check [ISSUES.md](./ISSUES.md) for known issues or to report a new one.
