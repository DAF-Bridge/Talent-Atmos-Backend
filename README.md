# Talent-Atmos-Backend

# Entity-Relationship Diagram

This ER diagram illustrates the relationships between various entities such as Organizations, Users, Events, and their related attributes.

``
Organization
- ID
- image
- logo
- name
- title
- Info **(have more detail)
- address
- email
- phone
- Organization Topic ID

Organization social media
- ID 
- Org ID
- media name => Website Facebook Instagram TikTok Youtube  
- text
- link

User Register Organization
- ID
- Org ID
- user ID
- Reg at

Organization opening job
- ID 
- Org ID
- title
- detail 
- scope
-  work at
- period
- work per day
- Qualifications 
- Benefits
- quantity
- create_at

User
-ID
- Username
- First Name
- last name
- phone number 
- email
- password
- image
- Job
- profile title
- list skill
- list language
-  list Education Background
-  list Working Experience

Organization Topic
- ID
- name

Event Topic
- ID
- name

Organization Category
- ID
- Organization Topic ID
- name

Event Category
- ID
- Event Topic ID
- name

Event
- ID
- Organization ID
- title
- image
- Date ** 
- time ** 
- location
- Event Category ID 
- detail
- Organization message

Event Ticket
- ID 
- event ID
- name
- Date ** 
- time ** 
- detail
- price
- quantity

Register Ticket
- ID
- Event Ticket ID
- User ID
- First Name
- last name
- phone number 
- email
- QR code
``


```mermaid
erDiagram

    %% Organization related entities
    Organization {
        int ID
        string image
        string logo
        string name
        string title
        string Info
        string address
        string email
        string phone
        int Organization_Topic_ID
    }

    Organization_social_media {
        int ID
        int Organization_ID
        string media_name
        string text
        string link
    }

    Organization_opening_job {
        int ID
        int Organization_ID
        string title
        string detail
        string scope
        string work_at
        string period
        string work_per_day
        string Qualifications
        string Benefits
        int Quantity
        string create_at
    }

    Organization_Topic {
        int ID
        string name
    }

    Organization_Category {
        int ID
        int Organization_Topic_ID
        string name
    }

    %% User related entities
    User {
        int ID
        string username
        string first_name
        string last_name
        string phone_number
        string email
        string password
        string image
        string job
        string profile_title
        string list_skill
        string list_language
        string list_education_background
        string list_working_experience
    }

    User_Register_Organization {
        int ID
        int Organization_ID
        int User_ID
        string Register_at
    }

    %% Event related entities
    Event {
        int ID
        int Organization_ID
        string title
        string image
        string date
        string time
        string location
        int Event_Category_ID
        string detail
        string Organization_message
    }

    Event_Ticket {
        int ID
        int event_ID
        string name
        string date
        string time
        string detail
        float price
        int quantity
    }

    Register_Ticket {
        int ID
        int Event_Ticket_ID
        int User_ID
        string first_name
        string last_name
        string phone_number
        string email
        string qr_code
    }

    Event_Topic {
        int ID
        string name
    }

    Event_Category {
        int ID
        int Event_Topic_ID
        string name
    }

    %% Relationships
    Organization ||--o{ Organization_social_media : "has"
    Organization ||--o{ Organization_opening_job : "offers"
    Organization ||--o{ Event : "hosts"
    Organization ||--o{ Organization_Category : "belongs to"
    Organization ||--o{ User_Register_Organization : "has registrations"
    
    User ||--o{ User_Register_Organization : "registers"
    User ||--o{ Register_Ticket : "registers for"

    Event ||--o{ Event_Ticket : "sells"
    Event ||--o{ Event_Category : "belongs to"

    Event_Ticket ||--o{ Register_Ticket : "issued for"

    Organization_Topic ||--o{ Organization_Category : "categorizes"
    Event_Topic ||--o{ Event_Category : "categorizes"

```
