function Message(type,message,duration){
    const msgHTML = `
<div class="alert alert-${type} alert-dismissible fade show" role="alert"
style="position: fixed; top: 20px; left: 50%; transform: translateX(-50%); min-width: 200px; text-align: center;">
  ${message}
  <button id="alert-close" type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
</div>`
    const msgDiv = document.createElement('div')
    msgDiv.innerHTML = msgHTML
    document.querySelector('body').appendChild(msgDiv)
    if (duration!==0){
        setTimeout(()=>{
            // document.querySelector('#alert-close').click()
            msgDiv.remove()
        },duration)
    }
}