// Données d'exemple pour simuler une API
const charactersData = [
    { id: 1, name: "Rick Sanchez", type: "Personnage", image: "/api/placeholder/250/200" },
    { id: 2, name: "Morty Smith", type: "Personnage", image: "/api/placeholder/250/200" },
    { id: 3, name: "Summer Smith", type: "Personnage", image: "/api/placeholder/250/200" },
    { id: 4, name: "Beth Smith", type: "Personnage", image: "/api/placeholder/250/200" },
    { id: 5, name: "Jerry Smith", type: "Personnage", image: "/api/placeholder/250/200" },
    { id: 6, name: "Abadango Alien", type: "Personnage", image: "/api/placeholder/250/200" }
];

const episodesData = [
    { id: 101, name: "Pilot", type: "Épisode", image: "/api/placeholder/250/200" },
    { id: 102, name: "Lawnmower Dog", type: "Épisode", image: "/api/placeholder/250/200" }
];

const locationsData = [
    { id: 201, name: "Earth (C-137)", type: "Lieu", image: "/api/placeholder/250/200" },
    { id: 202, name: "Citadel of Ricks", type: "Lieu", image: "/api/placeholder/250/200" }
];

// Exemple initial de favoris (simulation de stockage local)
let favorites = JSON.parse(localStorage.getItem('favorites')) || [];

// Fonction pour initialiser la page
function initPage() {
    renderFavorites();
    renderAddFavorites();
    renderPagination(1, 6, 8); // Page 1, 6 éléments par page, 8 éléments au total
    
    // Écouteurs d'événements
    document.getElementById('search-input').addEventListener('input', handleSearch);
    document.getElementById('browse-button').addEventListener('click', scrollToAddSection);
    document.querySelector('.toast-close').addEventListener('click', hideToast);
}

// Afficher les favoris
function renderFavorites() {
    const container = document.getElementById('favorites-container');
    const emptyState = document.getElementById('empty-state');
    
    if (favorites.length === 0) {
        container.style.display = 'none';
        emptyState.style.display = 'block';
        return;
    }
    
    container.style.display = 'grid';
    emptyState.style.display = 'none';
    
    container.innerHTML = '';
    favorites.forEach(item => {
        const card = createFavoriteCard(item, true);
        container.appendChild(card);
    });
}

// Afficher les éléments à ajouter aux favoris
function renderAddFavorites() {
    const container = document.getElementById('add-favorites-container');
    container.innerHTML = '';
    
    // Combinaison de toutes les données
    const allData = [...charactersData, ...episodesData, ...locationsData];
    
    // Filtrer les éléments qui ne sont pas déjà dans les favoris
    const notInFavorites = allData.filter(item => 
        !favorites.some(fav => fav.id === item.id && fav.type === item.type)
    );
    
    // Afficher les premiers éléments (pagination simple)
    const itemsToShow = notInFavorites.slice(0, 6);
    
    itemsToShow.forEach(item => {
        const card = createFavoriteCard(item, false);
        container.appendChild(card);
    });
}

// Créer une carte pour un élément
function createFavoriteCard(item, isFavorite) {
    const card = document.createElement('div');
    card.className = 'favorite-card';
    card.dataset.id = item.id;
    card.dataset.type = item.type;
    
    card.innerHTML = `
        <img src="${item.image}" alt="${item.name}">
        <div class="favorite-actions">
            ${isFavorite ? 
            `<button class="action-button remove-button" title="Retirer des favoris" onclick="removeFavorite(${item.id}, '${item.type}')">✕</button>` :
            `<button class="action-button add-button" title="Ajouter aux favoris" onclick="addFavorite(${item.id}, '${item.type}', '${item.name}', '${item.image}')">+</button>`}
        </div>
        <div class="favorite-info">
            <h3 class="favorite-name">${item.name}</h3>
            <p class="favorite-type">${item.type}</p>
            <div class="favorite-details">
                <button class="view-button" onclick="viewDetails(${item.id}, '${item.type}')">Voir détails</button>
            </div>
        </div>
    `;
    
    return card;
}

// Ajouter un élément aux favoris
function addFavorite(id, type, name, image) {
    const newFavorite = { id, type, name, image };
    
    // Vérifier si l'élément existe déjà dans les favoris
    if (!favorites.some(item => item.id === id && item.type === type)) {
        favorites.push(newFavorite);
        localStorage.setItem('favorites', JSON.stringify(favorites));
        
        // Mettre à jour l'affichage
        renderFavorites();
        renderAddFavorites();
        
        // Montrer la notification
        showToast('success', `${name} a été ajouté à vos favoris`);
    }
}

// Supprimer un élément des favoris
function removeFavorite(id, type) {
    const removedItem = favorites.find(item => item.id === id && item.type === type);
    favorites = favorites.filter(item => !(item.id === id && item.type === type));
    localStorage.setItem('favorites', JSON.stringify(favorites));
    
    // Mettre à jour l'affichage
    renderFavorites();
    renderAddFavorites();
    
    // Montrer la notification
    showToast('error', `${removedItem.name} a été retiré de vos favoris`);
}

// Afficher une notification toast
function showToast(type, message) {
    const toast = document.getElementById('toast');
    const toastMessage = document.querySelector('.toast-message');
    const toastIcon = document.querySelector('.toast-icon');
    
    toast.className = `toast ${type}`;
    toastMessage.textContent = message;
    toastIcon.textContent = type === 'success' ? '✓' : '✕';
    
    toast.classList.add('show');
    
    // Cacher la notification après 3 secondes
    setTimeout(() => {
        hideToast();
    }, 3000);
}

// Cacher la notification toast
function hideToast() {
    const toast = document.getElementById('toast');
    toast.classList.remove('show');
}

// Chercher dans les favoris
function handleSearch(e) {
    const searchTerm = e.target.value.toLowerCase();
    const container = document.getElementById('favorites-container');
    const emptyState = document.getElementById('empty-state');
    
    if (favorites.length === 0) {
        return;
    }
    
    const filteredFavorites = favorites.filter(item => 
        item.name.toLowerCase().includes(searchTerm) || 
        item.type.toLowerCase().includes(searchTerm)
    );
    
    if (filteredFavorites.length === 0) {
        container.style.display = 'none';
        emptyState.style.display = 'block';
        document.querySelector('.empty-state h3').textContent = 'Aucun résultat trouvé';
        document.querySelector('.empty-state p').textContent = 'Essayez avec un autre terme de recherche';
    } else {
        container.style.display = 'grid';
        emptyState.style.display = 'none';
        
        container.innerHTML = '';
        filteredFavorites.forEach(item => {
            const card = createFavoriteCard(item, true);
            container.appendChild(card);
        });
    }
}

// Voir les détails d'un élément
function viewDetails(id, type) {
    // Rediriger vers la page de détails
    window.location.href = `/details?id=${id}&type=${type}`;
}

// Faire défiler jusqu'à la section pour ajouter des favoris
function scrollToAddSection() {
    document.querySelector('.add-favorites-section').scrollIntoView({
        behavior: 'smooth'
    });
}

// Créer la pagination
function renderPagination(currentPage, itemsPerPage, totalItems) {
    const paginationContainer = document.getElementById('pagination');
    const totalPages = Math.ceil(totalItems / itemsPerPage);
    
    paginationContainer.innerHTML = '';
    
    for (let i = 1; i <= totalPages; i++) {
        const button = document.createElement('button');
        button.textContent = i;
        if (i === currentPage) {
            button.classList.add('active');
        }
        
        button.addEventListener('click', () => {
            // Logique de changement de page ici
            renderPagination(i, itemsPerPage, totalItems);
            // Mettre à jour les éléments affichés (non implémenté pour simplifier)
        });
        
        paginationContainer.appendChild(button);
    }
}

// Initialiser la page au chargement
document.addEventListener('DOMContentLoaded', initPage);