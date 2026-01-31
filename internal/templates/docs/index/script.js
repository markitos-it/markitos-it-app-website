// State
let selectedCategory = 'All';
let searchQuery = '';

// Elements
const categoryFilters = document.getElementById('categoryFilters');
const searchInput = document.getElementById('searchInput');
const docsGrid = document.getElementById('docsGrid');
const noResults = document.getElementById('noResults');

if (docsGrid) {
    docsGrid.addEventListener('keydown', (e) => {
        const card = e.target.closest('.doc-card');
        if (!card) return;

        if (e.key === 'Enter' || e.key === ' ') {
            e.preventDefault();
            viewDocument(card.dataset.id);
        }
    });
}

// Category Filter
if (categoryFilters) {
    categoryFilters.addEventListener('click', (e) => {
        if (e.target.classList.contains('filter-pill')) {
            // Update active state
            categoryFilters.querySelectorAll('.filter-pill').forEach(pill => {
                pill.classList.remove('active');
            });
            e.target.classList.add('active');

            // Update selected category
            selectedCategory = e.target.dataset.category;

            // Filter documents
            filterDocuments();
        }
    });
}

// Search Input
if (searchInput) {
    searchInput.addEventListener('input', (e) => {
        searchQuery = e.target.value.toLowerCase();
        filterDocuments();
    });
}

// Filter Documents Function
function filterDocuments() {
    const cards = docsGrid.querySelectorAll('.doc-card');
    let visibleCount = 0;

    cards.forEach(card => {
        const category = card.dataset.category;
        const title = card.dataset.title.toLowerCase();
        const description = card.dataset.description.toLowerCase();
        const tags = card.dataset.tags.toLowerCase();

        // Check category filter
        const categoryMatch = selectedCategory === 'All' || category === selectedCategory;

        // Check search filter
        const searchMatch = searchQuery === '' ||
            title.includes(searchQuery) ||
            description.includes(searchQuery) ||
            tags.includes(searchQuery);

        if (categoryMatch && searchMatch) {
            card.style.display = 'flex';
            visibleCount++;
        } else {
            card.style.display = 'none';
        }
    });

    // Show/hide no results message
    if (visibleCount === 0) {
        noResults.style.display = 'block';
        docsGrid.style.display = 'none';
    } else {
        noResults.style.display = 'none';
        docsGrid.style.display = 'grid';
    }
}

// View Document Function - Navigate to document page
function viewDocument(docId) {
    window.location.href = `/docs/${docId}`;
}

// Initialize
filterDocuments();