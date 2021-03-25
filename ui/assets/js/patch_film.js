// Узнаём название таблицы из урла
//const link = document.location.href.split('/')
//const table = link[link.length - 1].split('?')[0]
const url = `http://${HOST}/post/update/${table}`
// const temp = document.location.href.split('=')
// let ID = temp[temp.length -1]
// const form = document.querySelector('.fields__wrapper')

// Ajax запрос
function patchRecord(data){   
    load(true) 
    fetch(url, {
        method: 'PATCH',
        body: data
    }).then( () => {
        document.location.href = `../table/${table}/`
        setTimeout(() => load(false), 1000)
    })
}

document.querySelector('.save__btn').addEventListener('click', e => {
    form.addEventListener('submit', e => {
        e.preventDefault()    
        let data = new FormData(form)    
        data.set('id', ID)
        patchRecord(data)
    })
})