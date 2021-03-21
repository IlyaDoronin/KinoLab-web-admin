const ACTOR_CREATE = "ACTOR-CREATE"
const AUTHOR_CREATE = "AUTHOR-CREATE"
const FILM_ACTOR_CREATE = "FILM-ACTOR-CREATE"
const FILM_AUTHOR_CREATE = "FILM-AUTHOR-CREATE"
const FILM_COMMENTS_CREATE = "FILM-COMMENTS-CREATE"
const FILM_GENRE_CREATE = "FILM-GENRE-CREATE"
const FILM_CREATE = "FILM-CREATE"
const GENRE_CREATE = "GENRE-CREATE"

const addButton = document.getElementById("add")

addButton.addEventListener("click", getURLTable)

function getURLTable(e) {

    e.preventDefault()

    let correntURL = document.location.href
    let reg = new RegExp("http?:\/\/.+:[0-9]+\/.+\/")
    let result = correntURL.replace(reg, '')

    document.location.href = `/create/${result}`
}

function adminFieldsReducer(action) {
    switch(action.type) {
        case SET_SELECTED_TABLE: {
            state.selected_table = action.table
            console.log(state.selected_table)
        }
    }
}