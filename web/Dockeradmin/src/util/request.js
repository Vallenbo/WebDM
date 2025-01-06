import {ElMessage} from 'element-plus'

export function get(url) {
    return fetch(url)
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return response.json();
        })
        .catch(error => {
            console.log(`Fetch error: ${error}`);
        });
}

export function post(url, data) {
    return fetch(url, {
        method: 'POST', body: JSON.stringify(data), headers: {
            'Content-Type': 'application/json'
        }
    })
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return response.json();
        })
        .catch(error => {
            ElMessage.error('操作失败.')
            console.log(`Fetch error: ${error}`);
            throw error;
        });
}

export function postRaw(url, data) {
    return fetch(url, {
        method: 'POST', body: JSON.stringify(data), headers: {
            'Content-Type': 'application/json'
        }
    })
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return response;
        })
        .catch(error => {
            ElMessage.error('操作失败.')
            console.log(`Fetch error: ${error}`);
            throw error;
        });
}