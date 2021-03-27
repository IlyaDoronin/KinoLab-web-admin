// function $(el){
//     return document.querySelector(el)
// }
// // Грузить через go
// const HOST = '26.118.43.71:5500'

// const link = document.location.href.split('/')
// const table = link[link.length - 1].split('?')[0]
// const recordsList = $('.objects__items')

function getAllFields(){  
    URL = `http://${HOST}/get/table/${table}/all`
    // Код для починки объёбаного api
    if (table == 'actor')
        URL = `http://${HOST}/get/table/actors/all`
    if (table == 'author')
        URL = `http://${HOST}/get/table/authors/all`
    if (table == 'film')
        URL = `http://${HOST}/get/table/films/all`
    if (table == 'film_actor')
        URL = `http://${HOST}/get/table/films_actors/all`
    if (table == 'film_author')
        URL = `http://${HOST}/get/table/films_authors/all`
    if (table == 'film_genre')
        URL = `http://${HOST}/get/table/films_genres/all`
    if (table == 'genre')
        URL = `http://${HOST}/get/table/genres/all`       
    fetch(URL, { 
        method: 'GET',
        headers: { 'Content-Type': 'application/json' }
    }).then(response => {
        return response.json()
    }).then( fields => {   
        // удаление старых записей
        recordsList.innerHTML = ''                    
        
        if (table == 'actor'){
            fields.actors.map( field => generateFields(field))
        }
        if (table == 'author'){
            fields.authors.map( field => generateFields(field))
        }
        if (table == 'film'){
            fields.films.map( field => generateFields(field))
        }
        if (table == 'film_actor'){
            fields.films_actors.map( field => generateFields(field))
        }
        if (table == 'film_author'){
            fields.films_authors.map( field => generateFields(field))
        }
        if (table == 'film_comments'){
            fields.film_comments.map( field => generateFields(field))
        }
        if (table == 'film_genre'){
            fields.films_genres.map( field => generateFields(field))
        }
        if (table == 'genre'){
            fields.genres.map( field => generateFields(field))
        }
        
    })
}
// генерация HTML кода
function generateFields(field){
    let record = document.createElement('li')
    record.classList.add('objects__item')
    // Создание разных записей в зависимости от таблицы
    if (table == 'actor'){
        record.innerHTML = `<a href="/edit/actor?id=${field.ID}" class="object"><p>${field.FName} ${field.LName} id(${field.ID})</p></a>`
    }
    if (table == 'author'){
        record.innerHTML = `<a href="/edit/author?id=${field.ID}" class="object"><p>${field.FName} ${field.LName} id(${field.ID})</p></a>`
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


$('.objects__show').addEventListener('click', function(e){
    e.preventDefault()    
    getAllFields()
})
