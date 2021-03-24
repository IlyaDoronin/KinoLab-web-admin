// Узнаём название таблицы из урла
const link = document.location.href.split('/')
const table = link[link.length - 1].split('?')[0]
const URL = `http://${HOST}/post/create/${table}`

const form = document.querySelector('.fields__wrapper')

// Ajax запрос новых комментов
function sendComment(data){    
    fetch(URL, {
        method: 'POST',
        body: data
    }).then( () => document.location.href = `../table/${table}/`)

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
