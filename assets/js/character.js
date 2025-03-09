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