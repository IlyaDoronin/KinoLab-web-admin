const url = `http://${HOST}/post/update/film_comment`
// const form = document.querySelector('.fields__wrapper')
// const temp = document.location.href.split('=')
// let ID = temp[temp.length -1]

// Ajax запрос
function patchRecord(data){    
    fetch(url, {
        method: 'PATCH',
        body: data
    }).then( () => document.location.href = `../table/film_comments/`)
}


document.querySelector('.save__btn').addEventListener('click', e => {
    form.addEventListener('submit', e => {
        e.preventDefault()    
        let data = new FormData(form)    
        data.set('id', ID)
        patchRecord(data)
    })
})