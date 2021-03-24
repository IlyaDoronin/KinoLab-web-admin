const URL = `http://${HOST}/post/create/film_comment`

const form = document.querySelector('.fields__wrapper')

// Ajax запрос новых комментов
function sendComment(data){    
    fetch(URL, {
        method: 'POST',
        body: data
    })
    .then( () => document.location.href = `../table/film_comments/`)
}

function checkFields(){
    let checked = true
    document.querySelectorAll('input').forEach(input => {
        if (input.value == '')
            checked = false
    })
    return checked
}

form.addEventListener('submit', e => {
    e.preventDefault()
    if( checkFields() ){
        const data = new FormData(form)
        sendComment(data)    
    }
})