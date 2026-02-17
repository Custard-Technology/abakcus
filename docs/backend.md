I. Core Infrastructure & Authentication (MongoDB Focused)

MongoDB Schema Design:
    Ticket ID: BE-002
    Type: Technical (Database)
    Priority: High

    1. Businesses Collection:
        business_id:  String (UUID – recommended for unique identification)
        name: String (Required - Business Name)
        address: String (Required – Address line 1, Address line 2 - consider splitting into separate fields if needed)
        city: String (Optional – City)
        state: String (Optional – State/Province)
        zip_code: String (Optional – Zip Code / Postal Code)
        country: String (Optional - Country)
        contact_email: String (Optional – Email for contact)
        phone_number: String (Optional – Phone Number)
        website_url: String (Optional - Website URL)
        logo_url: String (Optional – URL to logo image)
        created_at: Date (Timestamp - Automatically populated on creation)
        updated_at: Date (Timestamp – Automatically updated on modification)

        Indexing:
        business_id: Index (Unique – Primary Key) - This is automatically created as the primary key.
        name: Index (for searching businesses by name) – Consider a text index if you need full-text search capabilities.
        city: Index (for filtering businesses by location)
