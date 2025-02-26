document.addEventListener('DOMContentLoaded', function() {
    // Handle favorite button click
    const favoriteBtn = document.querySelector('.favorite-btn');
    if (favoriteBtn) {
        favoriteBtn.addEventListener('click', function() {
            const characterId = this.getAttribute('data-id');
            toggleFavorite(characterId);
        });
        
        // Check initial favorite status
        const characterId = favoriteBtn.getAttribute('data-id');
        updateFavoriteButtonState(characterId);
    }
    
    // Add scroll animation for episodes
    const episodeCards = document.querySelectorAll('.episode-card');
    let delay = 0;
    
    episodeCards.forEach(card => {
        card.style.opacity = '0';
        card.style.transform = 'translateY(20px)';
        card.style.transition = 'opacity 0.5s ease, transform 0.5s ease';
        
        setTimeout(() => {
            card.style.opacity = '1';
            card.style.transform = 'translateY(0)';
        }, delay);
        
        delay += 50; // Stagger effect
    });
});

// Toggle character favorite status in localStorage
function toggleFavorite(characterId) {
    let favorites = JSON.parse(localStorage.getItem('rmFavorites') || '[]');
    const index = favorites.indexOf(characterId);
    
    if (index === -1) {
        // Add to favorites
        favorites.push(characterId);
        showToast('Personnage ajouté aux favoris !');
    } else {
        // Remove from favorites
        favorites.splice(index, 1);
        showToast('Personnage retiré des favoris.');
    }
    
    localStorage.setItem('rmFavorites', JSON.stringify(favorites));
    updateFavoriteButtonState(characterId);
}

// Update favorite button appearance based on favorite status
function updateFavoriteButtonState(characterId) {
    const favoriteBtn = document.querySelector('.favorite-btn');
    const favoriteBtnSvg = favoriteBtn.querySelector('svg path');
    let favorites = JSON.parse(localStorage.getItem('rmFavorites') || '[]');
    
    if (favorites.includes(characterId)) {
        favoriteBtn.classList.add('is-favorite');
        favoriteBtn.innerHTML = `
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="24" height="24">
                <path d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z" fill="#ff6b6b"/>
            </svg>
            Retiré des favoris
        `;
    } else {
        favoriteBtn.classList.remove('is-favorite');
        favoriteBtn.innerHTML = `
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="24" height="24">
                <path d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z"/>
            </svg>
            Ajouter aux favoris
        `;
    }
}

// Show toast notification
function showToast(message) {
    // Create toast element if it doesn't exist
    let toast = document.getElementById('toast-notification');
    
    if (!toast) {
        toast = document.createElement('div');
        toast.id = 'toast-notification';
        document.body.appendChild(toast);
        
        // Add toast styles
        toast.style.position = 'fixed';
        toast.style.bottom = '20px';
        toast.style.right = '20px';
        toast.style.backgroundColor = 'rgba(17, 176, 200, 0.9)';
        toast.style.color = 'white';
        toast.style.padding = '12px 24px';
        toast.style.borderRadius = '4px';
        toast.style.boxShadow = '0 4px 12px rgba(0, 0, 0, 0.2)';
        toast.style.zIndex = '1000';
        toast.style.transition = 'transform 0.3s ease, opacity 0.3s ease';
        toast.style.transform = 'translateY(100px)';
        toast.style.opacity = '0';
    }
    
    // Set message and show toast
    toast.textContent = message;
    toast.style.transform = 'translateY(0)';
    toast.style.opacity = '1';
    
    // Hide toast after 3 seconds
    setTimeout(() => {
        toast.style.transform = 'translateY(100px)';
        toast.style.opacity = '0';
    }, 3000);
}