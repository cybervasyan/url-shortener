import React, {useState, useEffect} from "react";
import axios from 'axios';
import './Home.css';

function Home() {
    const [message, setMessage] = useState("");
    const [url, setUrl] = useState("");
    const [userId, setUserId] = useState('e0dba740-fc4b-4977-872c-d360239e6b10');
    const [shortUrl, setShortUrl] = useState("");
    const [error, setError] = useState("");

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

    const handleSubmit = async (e) => {
        e.preventDefault();
        setError('');
        setShortUrl('');

        if (!url) {
            setError('Пожалуйста, введите URL пользователя.');
            return;
        }

        try {
            const response = await axios.post('http://localhost:1488/create-short-url', {
                longUrl: url,
                userId: userId
            });
            if (response.data.shortUrl) {
                setShortUrl(response.data.shortUrl);
                setError('')
            } else {
                throw new Error('Сервер вернул ошибку');
            }
        } catch (err) {
            setError(err.response?.data?.message || 'Ошибка при создании сокращённого URL.');
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
            {shortUrl && <p>Сокращённый URL: <a href={shortUrl}>{shortUrl}</a></p>}
            {error && <p>{error}</p>}
        </div>
    );

}

export default Home;
