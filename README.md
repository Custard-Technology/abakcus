![](https://raw.githubusercontent.com/Custard-Technology/custard-app/refs/heads/main/frontend/public/hom2.png)

# Abakcus: Loyalty Points & Voucher Program

**Abakcus** is a simple application designed to help small businesses effortlessly manage loyalty points and vouchers.

## Our Mission

To empower small businesses by simplifying loyalty management and boosting customer retention.

## Our Values

*   **Simplicity**: Easy-to-use and intuitive for everyone.
*   **Affordability**: Cost-effective solutions for small businesses.
*   **Reliability**: A dependable platform you can count on.
*   **Empowerment**: Giving businesses the tools to grow.

## Target Audience

*   **Primary Users**: Small business owners (retailers, coffee shops, salons, etc.).
*   **Secondary Users**: Loyal customers of these small businesses.

## Branding

### Color Palette

*   **Primary**: Blue (for trust and reliability) or Green (for growth and success).
*   **Accent**: Yellow/Gold (for optimism and rewards).
*   **Neutral**: White/Light Gray (for simplicity and clarity).

### UI/UX

The user interface should be clean, modern, and aligned with our branding. Navigation and user flows are designed to be simple and intuitive.

## Application Flow

### Website Flow

1.  **Landing Page**: A simple page with a hero section explaining the product and a footer with basic legal information and a sign-in button.
2.  **Sign-in Page**: Users sign in with their email. The system checks for referrals. New users are onboarded; existing users are sent to their dashboard.
3.  **Onboarding**: A multi-step process to collect basic business and personal details.
4.  **Dashboard**: The main hub for businesses, featuring a side navigation menu for accessing different features.

### Dashboard Navigation

*   **Home**: Basic analytics dashboard.
*   **Loyalty**: Create and manage loyalty points.
*   **Cards**: Manage customer cards.
*   **Referrals**: Track customer referrals.
*   **Customer Directory**: View and manage customer information.
*   **Staff**: Manage staff access and permissions.
*   **Settings**: Configure business settings.

## Technical Overview

### Backend Flow

The backend is designed as a set of RESTful APIs to handle all business logic.

*   **Authentication**: Handles user sign-in, registration, and referral checks.
*   **Onboarding**: APIs for updating user and business details.
*   **Dashboard**: Endpoints for fetching analytics.
*   **Loyalty Points & Cards**: APIs for creating, assigning, and redeeming points.
*   **Customer & Staff Management**: Endpoints for managing customers and staff members.

### Database Design

The database schema is designed to be scalable and support multi-tenancy.

*   **Core Tables**: `Users`, `Businesses`, `LoyaltyPoints`, `Cards`, `Referrals`, `Customers`, `Staff`.

### Smart Contracts

The `contract` directory contains the Solidity smart contracts for the loyalty program.

*   `TokenFactory.sol`: A factory contract to deploy new ERC20-compliant loyalty point tokens.
*   `FactoryToken.sol`: An ERC20 token contract that is created by the factory.

## Project Structure

```
/contract
    ├── artifacts/      # Contract artifacts
    ├── cache/          # Hardhat cache
    ├── contracts/      # Solidity smart contracts
    ├── scripts/        # Deployment scripts
    └── test/           # Test files
```

```
/frontend
    ├── components/     # Reusable UI components
    ├── hooks/          # Custom React hooks
    ├── lib/            # Library functions
    ├── pages/          # Next.js pages and API routes
    ├── public/         # Static assets
    ├── store/          # State management
    └── styles/         # Global styles
```

## Getting Started

### Prerequisites

*   Node.js (v18 or later)
*   npm (v9 or later)

```bash
npm install npm@latest -g
```

### Environment Variables

Create a `.env` file in both the `contract` and `frontend` directories and add the required environment variables.

```
# contract/.env
RPC_URL=
PUBLIC_KEY=
PRIVATE_KEY=

# frontend/.env
MONGODB_URI=
RESEND_API=
FACTORY_CONTRACT_ADDRESS=
ENV=
LINK=
DATABASE=
```

### Contract Deployment

```bash
cd contract
npm install
npx hardhat deploy
npm run deploy:primodial
```

### Frontend Development

```bash
cd frontend
npm install
npm run dev
```

Open http://localhost:3000 to see the application.
