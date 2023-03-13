const { Menu, dialog } = require('electron')

const setMainMenu = (mainWindow) => {
  const template = [
    {
      label: 'mibuMarkDown',
      submenu: [
        { role: 'about' },
        { type: 'separator' },
        { role: 'services' },
        { type: 'separator' },
        { role: 'hide' },
        { role: 'hideOthers' },
        { role: 'unhide' },
        { type: 'separator' },
        { role: 'quit' }
      ]
    },
    {
      label: 'Themes',
      submenu: [
        {
          label: 'Light',
          click: () => {
            mainWindow.webContents.send('theme', 'light')
          }
        },
        {
          label: 'Dark',
          click: () => {
            mainWindow.webContents.send('theme', 'dark')
          }
        }
      ]
    },
    {
      label: 'View',
      submenu: [
        { role: 'reload' },
        { role: 'forceReload' },
        { role: 'toggleDevTools' },
        { type: 'separator' },
        { role: 'resetZoom' },
        { role: 'zoomIn' },
        { role: 'zoomOut' },
        { type: 'separator' },
        { role: 'togglefullscreen' }
      ]
    },
    {
      label: 'Edit',
      submenu: [
        {
          label: 'Abrir archivo',
          click: () => {
            dialog.showOpenDialogSync(mainWindow, {
              filters: {
                name: 'Markdown',
                extentions: ['md']
              },
              title: 'Selecciona tu archivo markdown',
              defaultPath: '~/Desktop',
              properties: ['openFile', 'openDirectory']
            }).then(result => {
              console.log({ lalalalala: result })
            }).catch(err => console.log({ lalalalala: err }))
          }
        }
      ]
    }
  ]

  const menu = Menu.buildFromTemplate(template)
  Menu.setApplicationMenu(menu)
}

module.exports = {
  setMainMenu
}
