const searchBar = document.querySelector('.search-bar');
if (searchBar) {
    searchBar.addEventListener('input', (e) => {
        const term = e.target.value.toLowerCase();
        const cards = document.querySelectorAll('.package-card');
        
        cards.forEach(card => {
            const title = card.querySelector('h3').innerText.toLowerCase();
            card.style.display = title.includes(term) ? 'block' : 'none';
        });
    });
}

const filterCheckboxes = document.querySelectorAll('.home-filters input[type="checkbox"]');
filterCheckboxes.forEach(checkbox => {
    checkbox.addEventListener('change', () => {
        console.log('Filter changed:', checkbox.parentElement.textContent.trim());
    });
});