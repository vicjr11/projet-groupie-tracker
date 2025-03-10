document.addEventListener('DOMContentLoaded', function() {
    // Récupérer les favoris du localStorage
    const favorites = JSON.parse(localStorage.getItem('rickMortyFavorites')) || [];
    const favoritesContainer = document.querySelector('.favorites-grid');
    
    if (favorites.length === 0) {
        // Aucun favori
        favoritesContainer.innerHTML = '<p class="no-favorites">Vous n\'avez pas encore de favoris.</p>';
    } else {
        // Afficher les favoris
        favorites.forEach(character => {
            const characterCard = document.createElement('article');
            characterCard.className = 'character-card fade-in';
            
            characterCard.innerHTML = `
                <img src="${character.image}" alt="${character.name}">
                <div class="character-info">
                    <h3>${character.name}</h3>
                    <button class="remove-fav-btn" data-id="${character.id}">
                        🗑️ Retirer des favoris
                    </button>
                </div>
            `;
            
            favoritesContainer.appendChild(characterCard);
        });
        
        // Ajouter des écouteurs d'événements pour les boutons de suppression
        document.querySelectorAll('.remove-fav-btn').forEach(button => {
            button.addEventListener('click', function(e) {
                const characterId = e.currentTarget.getAttribute('data-id');
                
                // Retirer des favoris
                const updatedFavorites = favorites.filter(fav => fav.id !== characterId);
                localStorage.setItem('rickMortyFavorites', JSON.stringify(updatedFavorites));
                
                // Supprimer la carte du DOM
                e.currentTarget.closest('.character-card').remove();
                
                // Vérifier s'il reste des favoris
                if (updatedFavorites.length === 0) {
                    favoritesContainer.innerHTML = '<p class="no-favorites">Vous n\'avez pas encore de favoris.</p>';
                }
            });
        });
    }
});