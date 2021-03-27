function $(el){
    return document.querySelector(el)
}
// Грузить через go
const HOST = '26.118.43.71:5500'

const link = document.location.href.split('/')
const table = link[link.length - 1].split('?')[0]
const recordsList = $('.objects__items')
const paginationList = $('.object__pages-items')

function getFields(page){  
    URL = `http://${HOST}/get/table/${table}?page=${page}`
    // Код для починки объёбаного api
    if (table == 'actor')
        URL = `http://${HOST}/get/table/actors?page=${page}`
    if (table == 'author')
        URL = `http://${HOST}/get/table/authors?page=${page}`
    if (table == 'film')
        URL = `http://${HOST}/get/table/films?page=${page}`
    if (table == 'film_actor')
        URL = `http://${HOST}/get/table/films_actors?page=${page}`
    if (table == 'film_author')
        URL = `http://${HOST}/get/table/films_authors?page=${page}`
    if (table == 'film_genre')
        URL = `http://${HOST}/get/table/films_genres?page=${page}`
    if (table == 'genre')
        URL = `http://${HOST}/get/table/genres?page=${page}`        
    fetch(URL, { 
        method: 'GET',
        headers: { 'Content-Type': 'application/json' }
    }).then(response => {
        return response.json()
    }).then( fields => {   
        // удаление старых записей
        recordsList.innerHTML = ''                    
        
        if (table == 'actor'){
            fields.actors.map( field => generateField(field))
        }
        if (table == 'author'){
            fields.authors.map( field => generateField(field))
        }
        if (table == 'film'){
            fields.films.map( field => generateField(field))
        }
        if (table == 'film_actor'){
            fields.films_actors.map( field => generateField(field))
        }
        if (table == 'film_author'){
            fields.films_authors.map( field => generateField(field))
        }
        if (table == 'film_comments'){
            fields.film_comments.map( field => generateField(field))
        }
        if (table == 'film_genre'){
            fields.films_genres.map( field => generateField(field))
        }
        if (table == 'genre'){
            fields.genres.map( field => generateField(field))
        }
        
    })
}
// генерация HTML кода
function generateField(field){
    let record = document.createElement('li')
    record.classList.add('objects__item')
    // Создание разных записей в зависимости от таблицы
    if (table == 'actor'){
        record.innerHTML = `<a href="/edit/actor?id=${field.ID}" class="object"><p>${field.LName} ${field.FName} id(${field.ID})</p></a>`
    }
    if (table == 'author'){
        record.innerHTML = `<a href="/edit/author?id=${field.ID}" class="object"><p>${field.LName} ${field.FName} id(${field.ID})</p></a>`
    }
    if (table == 'film'){
        record.innerHTML = `<a href="/edit/film?id=${field.ID}" class="object"><p>${field.FilmName}  id(${field.ID})</p></a>`
    }
    if (table == 'film_actor'){
        record.innerHTML = `<a href="/edit/film_actor?id=${field.ID}" class="object"><p>${field.ActorName} film(${field.FilmName}) id(${field.ID})</p></a>`
    }
    if (table == 'film_author'){
        record.innerHTML = `<a href="/edit/film_author?id=${field.ID}" class="object"><p>${field.AuthorName} film(${field.FilmName}) id(${field.ID})</p></a>`
    }
    if (table == 'film_comments'){
        record.innerHTML = `<a href="/edit/film_comment?id=${field.ID}" class="object"><p>${field.Name} film_id(${field.FilmID}) id(${field.ID})</p></a>`
    }
    if (table == 'film_genre'){
        record.innerHTML = `<a href="/edit/film_genre?id=${field.ID}" class="object"><p>${field.GenreName} film(${field.FilmName}) id(${field.ID})</p></a>`
    }
    if (table == 'genre'){
        record.innerHTML = `<a href="/edit/genre?id=${field.ID}" class="object"><p>${field.GenreName} id(${field.ID})</p></a>`
    }
    recordsList.append(record)
}

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

let btns = document.querySelectorAll('.object__page').forEach(btn => {
    btn.addEventListener('click', function(e){
        e.preventDefault()    
        $('.object__page.active').classList.remove('active')
        this.classList.add('active')
        getFields(this.innerHTML)
    })
})  
