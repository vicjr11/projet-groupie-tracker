let currentPage = 1;

async function fetchCharacters(page = 1) {
    const response = await fetch(`https://rickandmortyapi.com/api/character?page=${page}`);
    const data = await response.json();
    return data;
}

function renderCharacters(characters) {
    const characterGrid = document.getElementById('characterGrid');
    characterGrid.innerHTML = '';
    characters.forEach(character => {
        const characterCard = document.createElement('div');
        characterCard.className = 'character-card';
        characterCard.innerHTML = `
            <img src="${character.image}" alt="${character.name}">
            <h2>${character.name}</h2>
        `;
        characterGrid.appendChild(characterCard);
    });
}

async function loadPage(page) {
    const data = await fetchCharacters(page);
    renderCharacters(data.results);
    document.getElementById('pageInfo').textContent = `Page ${page}`;
}

document.getElementById('prevPage').addEventListener('click', () => {
    if (currentPage > 1) {
        currentPage--;
        loadPage(currentPage);
    }
});

document.getElementById('nextPage').addEventListener('click', () => {
    currentPage++;
    loadPage(currentPage);
});

document.getElementById('search').addEventListener('input', async (e) => {
    const searchTerm = e.target.value.toLowerCase();
    const data = await fetchCharacters(currentPage);
    const filteredCharacters = data.results.filter(character => character.name.toLowerCase().includes(searchTerm));
    renderCharacters(filteredCharacters);
});

document.getElementById('statusFilter').addEventListener('change', async (e) => {
    const status = e.target.value;
    const data = await fetchCharacters(currentPage);
    const filteredCharacters = data.results.filter(character => status === '' || character.status === status);
    renderCharacters(filteredCharacters);
});

document.getElementById('speciesFilter').addEventListener('change', async (e) => {
    const species = e.target.value;
    const data = await fetchCharacters(currentPage);
    const filteredCharacters = data.results.filter(character => species === '' || character.species === species);
    renderCharacters(filteredCharacters);
});

loadPage(currentPage);



// Gestion des favoris
document.addEventListener('DOMContentLoaded', function() {
    // Récupérer les favoris du localStorage
    let favorites = JSON.parse(localStorage.getItem('rickMortyFavorites')) || [];
    
    // Mettre à jour l'apparence des boutons favoris
    updateFavoriteButtons();
    
    // Ajouter des écouteurs d'événements pour tous les boutons favoris
    document.querySelectorAll('.fav-btn').forEach(button => {
        button.addEventListener('click', toggleFavorite);
    });
    
    // Fonction pour basculer un personnage des favoris
    function toggleFavorite(e) {
        const button = e.currentTarget;
        const characterId = button.getAttribute('data-id');
        const characterName = button.getAttribute('data-name');
        const characterImage = button.getAttribute('data-image');
        
        // Vérifier si le personnage est déjà dans les favoris
        const index = favorites.findIndex(fav => fav.id === characterId);
        
        if (index === -1) {
            // Ajouter aux favoris
            favorites.push({
                id: characterId,
                name: characterName,
                image: characterImage
            });
            button.classList.add('active');
            button.textContent = '❤️ Retiré des favoris';
        } else {
            // Retirer des favoris
            favorites.splice(index, 1);
            button.classList.remove('active');
            button.textContent = '❤️ Favoris';
        }
        
        // Sauvegarder dans localStorage
        localStorage.setItem('rickMortyFavorites', JSON.stringify(favorites));
        
        // Mettre à jour l'apparence des boutons
        updateFavoriteButtons();
    }
    
    // Mettre à jour l'apparence des boutons favoris
    function updateFavoriteButtons() {
        document.querySelectorAll('.fav-btn').forEach(button => {
            const characterId = button.getAttribute('data-id');
            const isFavorite = favorites.some(fav => fav.id === characterId);
            
            if (isFavorite) {
                button.classList.add('active');
                button.textContent = '❤️ Retiré des favoris';
            } else {
                button.classList.remove('active');
                button.textContent = '❤️ Favoris';
            }
        });
    }
});