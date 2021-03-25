const loading = document.querySelector(".loading");

// Анимация загрузки
function load(start, delay = 0){
    setTimeout( () => {
        if (start){
            loading.style.transition = ".5s";
            loading.style.left = "0";                
        }
        else{      
            loading.style.left = "100vw";
            setTimeout( () => {
                loading.style.transition = "0s";
                loading.style.left = "-100vw";
            }, 500)
        }
    }, delay)
}