## GoHire 
<img width="467" height="426" alt="image" src="https://github.com/user-attachments/assets/18b25424-b172-4f5e-833e-801bc54bd45d" />

GoHire will be a project which will scrape hiring data from different websites available on the internet and will present to you on a single a web page or on your terminal. So instead of searching different websites for finding jobs and internships you only have to search on GoHire and it will categorically search different websites and scrape the relevant info.

### High level architecture overview

				  ┌─────────────────────────────┐
                  │         Presentation        │
                  │ (CLI or Web UI - optional)  │
                  └─────────────────────────────┘
                             │
                             ▼
                  ┌─────────────────────────────┐
                  │      Aggregation Layer      │
                  │  - Takes input & triggers   │
                  │    concurrent scrapers      │
                  └─────────────────────────────┘
                             │
    ┌────────────────────────┼────────────────────────┐
    ▼                        ▼                        ▼
	┌─────────────┐         ┌─────────────┐          ┌─────────────┐
	│  Indeed     │         │   Hirist    │          │  RemoteOK   │
	│ Scraper     │         │ Scraper     │          │ Scraper     │
	└─────────────┘         └─────────────┘          └─────────────┘
	│                        │                        │
	└─────────────┬──────────┴──────────┬─────────────┘
							▼                     ▼
		(Optional) Cache/Storage     (Optional)
	┌─────────────────────────────┐
	│   In-Memory / Redis / DB    │
	└─────────────────────────────┘

*Here is a notion file for the design of the project* - *https://www.notion.so/GoHire-248179a134d6801e911ae55c36caeec0?source=copy_link*
