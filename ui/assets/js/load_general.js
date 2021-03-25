function $(el){
    return document.querySelector(el)
}
// Грузить через go
const HOST = '26.118.43.71:5600'

const link = document.location.href.split('/')
const table = link[link.length - 1].split('?')[0]
const recordsList = $('.objects__items')
const paginationList = $('.object__pages-items')

function getFields(page){    
    URL = `http://${HOST}/get/table/${table}?page=${page}`
    fetch(URL, { 
        method: 'GET',
        headers: { 'Content-Type': 'application/json' }
    }).then(response => {
        return response.json()
    }).then( fields => {   
        // удаление старых записей
        recordsList.innerHTML = ''                       
        fields[table].map( field => generateField(field))
        // console.log(field);
    })
}

// генерация HTML кода
function generateField(field){
    let record = document.createElement('li')
    record.classList.add('objects__item')
    record.innerHTML = `<a href="/edit/${table}?id=2" class="object"><p>${field.LName} ${field.FName} id(${field.ID})</p></a>`
    recordsList.append(record)
}
document.querySelectorAll('.object__page').forEach( btn => {
    debugger
    btn.addEventListener('click', function(e) {        
        // e.preventDefault()
        console.log(1);
    
        // $('.object__page.active').classList.remove('active')
        // this.classList.add('active')
        // getFields(this.innerHTML)
    })
})
// Генерация кнопочек пагинации
function generatePagination(){
    paginationList.innerHTML = ''
    for (let i = 1; i <= pageCount; i++){
        let pag = document.createElement('button')
        pag.classList.add('object__page')
        pag.innerHTML = i
        paginationList.append(pag)
    }       
    $('.object__page').classList.add('active')
}
generatePagination()