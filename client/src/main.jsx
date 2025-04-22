import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.jsx'
import './App.css'; // make sure this line is present


createRoot(document.getElementById('root')).render(
  
    <App />
  
)
