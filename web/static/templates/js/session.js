(function () {
    // Функция для получения куки по имени
    function getCookie(name) {
      const value = `; ${document.cookie}`;
      const parts = value.split(`; ${name}=`);
      if (parts.length === 2) return parts.pop().split(';').shift();
    }
  
    const sessionID = getCookie('session_id');
  
    if (!sessionID) {
      // Если сессии нет — создаём новую
      fetch('/api/session', {
        method: 'POST',
        credentials: 'include', // 👈 важно! чтобы браузер принимал Set-Cookie
      }).then((res) => {
        if (!res.ok) {
          console.error('Не удалось создать сессию');
        }
      });
    }
  })();
  