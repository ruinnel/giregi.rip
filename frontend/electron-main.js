const { app, BrowserWindow } = require('electron');
const path = require('path');
const { execFile } = require('child_process');
const fs = require('fs');

let mainWindow;
let server = null;

function runServer() {
  // process.env.ELECTRONVUESITE = app.getAppPath();
  // process.env.ELECTRONVUESITE = process.env.ELECTRONVUESITE.replace('\\app.asar', '');
  const appDataPath = path.join(app.getPath('appData'), 'rip.giregi.app');
  process.env.APP_DATA_PATH = appDataPath;

  if (!fs.existsSync(appDataPath)) {
    fs.mkdirSync(appDataPath);
  }

  const configPath = path.join(__dirname, 'dist', '/config.yaml');
  console.log('app data path - ', appDataPath);
  console.log('app path - ', app.getAppPath());
  console.log('dir - ', __dirname);
  console.log('configPath - ', configPath);
  let executablePath = '';
  if (process.platform === 'win32') {
    executablePath = path.join(__dirname, 'dist', '/server.exe');
  } else if (process.platform === 'linux') {
    executablePath = path.join(__dirname, 'dist', '/server');
  } else if (process.platform === 'darwin') {
    executablePath = path.join(__dirname, 'dist', '/server');
  }

  server = execFile(executablePath, ['-config', configPath], {}, (err, stdout, stderr) => {
    if (err) {
      console.error('run server fail - ', err);
      throw err;
    }

    console.log('stdout - ', stdout.toString());
    console.log('stderr - ', stderr.toString());
  });
  server.stdout.on('data', (chunk) => { console.log(chunk.toString()); });
  server.stderr.on('data', (chunk) => { console.log(chunk.toString()); });
}

function createWindow() {
  mainWindow = new BrowserWindow({
    width: 1026,
    height: 768,
    webPreferences: {
      nodeIntegration: true,
      contextIsolation: true,
      enableRemoteModule: true,
      devTools: true,
    },
  });

  mainWindow.loadFile('./dist/index.html');

  mainWindow.on('closed', function () {
    mainWindow = null;
  });

  // callback for - target='_blank' link.
  mainWindow.webContents.on('new-window', function(e, url) {
    e.preventDefault();
    require('electron').shell.openExternal(url);
  });
}

app.once('ready', function () {
  if (server == null) {
    runServer();
  }
  createWindow();
  // mainWindow.webContents.openDevTools();
});
app.on('window-all-closed', function () {
  if (process.platform !== 'darwin') {
    app.quit();
  }
});

app.on('activate', function () {
  if (mainWindow === null) {
    createWindow();
  }
});
