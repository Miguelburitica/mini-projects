const $ = seletor => document.querySelector(seletor)

const $count = $('#count')
const $button = $('button')

$button.addEventListener('click', () => {
  const count = +$count.innerText
  $count.innerText = count + 1
})

window.electronAPI.onUpdateTheme((event, theme) => {
  console.log(' ------------------------> HOlaaa')
  const root = document.documentElement

  root.style.setProperty('--scheme', theme)
  console.log({ lalalalala: event, theme })
})
