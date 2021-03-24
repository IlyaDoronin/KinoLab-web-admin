function $(el){
    return document.querySelector(el)
}
// console.log(pageCount);
const HOST = '26.118.43.71:5600'
const table = 'actors'
const recordsList = $('.objects__items')


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
    })

}

// генерация HTML кода
function generateField(field){
    let record = document.createElement('li')
    record.classList.add('objects__item')
    record.innerHTML = `<a href="/edit/actor?id=2" class="object"><p>${field.LName} ${field.FName} id(${field.ID})</p></a>`
    recordsList.append(record)
}

document.querySelectorAll('.object__page').forEach( btn => {
    btn.addEventListener('click', function() {        
        $('.object__page.active').classList.remove('active')
        this.classList.add('active')
        getFields(this.innerHTML)
    })
})