
// Узнаём название таблицы из урла
const link = document.location.href.split('/')
const table = link[link.length - 1].split('?')[0]
const temp = document.location.href.split('=')
let ID = temp[temp.length -1]
const URL = `http://${HOST}/post/delete/${table}`
const form = document.querySelector('.fields__wrapper')


// Ajax запрос новых комментов
function sendComment(data){    
    fetch(URL, {
        method: 'DELETE',
        body: data
    }).then( () => document.location.href = `../table/${table}/`)
}
document.querySelector('.delete__btn').addEventListener('click', e => {
    form.addEventListener('submit', e => {            
        e.preventDefault()     
        let data = new FormData(form)
        data.set('id', ID)
        sendComment(data)
    })
})