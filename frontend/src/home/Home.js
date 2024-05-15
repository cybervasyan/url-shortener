import React, {useState, useEffect} from "react";
import axios from 'axios';
import './Home.css';
import QRCode from 'qrcode.react';


function Home() {
    const [message, setMessage] = useState("");
    const [url, setUrl] = useState("");
    const [userId, setUserId] = useState('e0dba740-fc4b-4977-872c-d360239e6b10');
    const [shortUrl, setShortUrl] = useState("");
    const [error, setError] = useState("");
    const [qrCode, setQrCode] = useState("");


    useEffect(() => {
        fetch('http://localhost:1488/')
            .then(response => response.json())
            .then(data => {
                setMessage(data.message);
            })
            .catch(error => {
                console.error('Ошибка при загрузке данных:', error);
                setMessage('Не удалось загрузить данные');
            });
    }, []);

    const isValidUrl = (url) => {
        const urlPattern = new RegExp('^(https?:\\/\\/)?' +
            '((([a-z\\d]([a-z\\d-]*[a-z\\d])*)\\.?)+[a-z]{2,}|' +
            '((\\d{1,3}\\.){3}\\d{1,3}))' +
            '(\\:\\d+)?(\\/[-a-z\\d%_.~+]*)*' +
            '(\\?[;&a-z\\d%_.~+=-]*)?' +
            '(\\#[-a-z\\d_]*)?$', 'i');
        return !!urlPattern.test(url);
    }

    const handleSubmit = async (e) => {
        e.preventDefault();
        setError('');
        setShortUrl('');

        if (!url) {
            setError('Пожалуйста, введите URL пользователя.');
            return;
        }

        if (!isValidUrl(url)) {
            setError('Пожалуйста, введите действительный URL.');
            return;
        }

        try {
            const response = await axios.post('http://localhost:1488/create-short-url', {
                longUrl: url,
                userId: userId
            });
            if (response.data.short_url) {
                setShortUrl(response.data.short_url);
                setQrCode(response.data.qr_code);
            } else {
                throw new Error('Сервер вернул ошибку');
            }
        } catch (err) {
            setError(err.response?.data?.error || 'Ошибка при создании сокращённого URL.');
        }
    };

    return (
        <div className="home">
            <h1>{message}</h1>
            <form onSubmit={handleSubmit} className="url-form">
                <input
                    type="text"
                    placeholder="Введите URL"
                    value={url}
                    onChange={(e) => setUrl(e.target.value)}
                    required
                />
                <button type="submit">Сократить URL</button>
            </form>
            {shortUrl && (
                <div>
                    <p>Сокращённый URL: <a href={shortUrl} target="_blank" rel="noopener noreferrer">{shortUrl}</a></p>
                    <QRCode value={shortUrl} />
                </div>
            )}
            {error && <p className="error">{error}</p>}
        </div>
    );

}

export default Home;
