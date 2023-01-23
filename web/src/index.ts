import './style.css'
const htmx = require('htmx.org')

import {library, dom} from '@fortawesome/fontawesome-svg-core'
import {faCamera} from '@fortawesome/free-solid-svg-icons'
import {faRss} from '@fortawesome/free-solid-svg-icons'
import {faTrashCan} from '@fortawesome/free-solid-svg-icons'
import {faEllipsis} from '@fortawesome/free-solid-svg-icons'
import {faPlus} from '@fortawesome/free-solid-svg-icons'
import {faFolderOpen} from '@fortawesome/free-regular-svg-icons'

library.add(faCamera, faRss, faFolderOpen, faTrashCan, faEllipsis, faPlus)

// Replace any existing <i> tags with <svg> and set up a MutationObserver to
// continue doing this as the DOM changes.
dom.watch()


const tabHandler = function (event) {
    document.querySelectorAll('.ch-tab').forEach((tab) => {
        tab.addEventListener('click', (event) => {
            const target = event.target
            if (!(target instanceof Element)) {
                return
            }

            // Activate the current tab
            document.querySelectorAll('.ch-tab').forEach((tab) => {
                tab.classList.remove('tab-active')
            })
            target.classList.add('tab-active')
        })
    })

    document.querySelectorAll('table input[type=checkbox]').forEach((checkbox) => {
        checkbox.addEventListener('change', (event) => {
            const n = document.querySelectorAll('table input[type=checkbox]:checked').length
            console.log(n)
            const button = document.querySelector('#delete-selected')
            console.log(button)
            if (n === 0) {
                button.classList.add('disabled')
            } else {
                button.classList.remove('disabled')
            }
        })
    })
};


htmx.onLoad(function (content) {
    content.querySelectorAll('.ch-tab').forEach((tab) => {
        tab.addEventListener('click', (event) => {
            const target = event.target
            if (!(target instanceof Element)) {
                return
            }

            // Activate the current tab
            content.querySelectorAll('.ch-tab').forEach((tab) => {
                tab.classList.remove('tab-active')
            })
            target.classList.add('tab-active')
        })
    })

    content.querySelectorAll('table input[type=checkbox]').forEach((checkbox) => {
        checkbox.addEventListener('change', (event) => {
            const n = content.querySelectorAll('table input[type=checkbox]:checked').length
            console.log(n)
            const button = content.querySelector('#delete-selected')
            console.log(button)
            if (n === 0) {
                button.classList.add('disabled')
            } else {
                button.classList.remove('disabled')
            }
        })
    })
})

// document.addEventListener('htmx:afterRequest', tabHandler)
// document.addEventListener('DOMContentLoaded', tabHandler)
