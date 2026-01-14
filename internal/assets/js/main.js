document.addEventListener('DOMContentLoaded', () => {
    console.log("Dashboard cargado correctamente para markitos.it");

    const cards = document.querySelectorAll('.card');

    cards.forEach(card => {
        card.addEventListener('click', () => {
            const title = card.querySelector('h3').innerText;
            alert(`Abriendo: ${title}`);
        });
    });
});
