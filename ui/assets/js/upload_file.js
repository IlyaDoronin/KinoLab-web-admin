
  var reader;
  var progress = document.querySelector('.percent');
  function abortRead() {
    reader.abort();
  }

  const pFilm = document.querySelector('#progress_film > div')
  const tFilm = document.querySelector('#progress_film > p')
  const pPoster = document.querySelector('#progress_poster > div')
  const tPoster = document.querySelector('#progress_poster > p')
  const pBanner = document.querySelector('#progress_banner > div')
  const tBanner = document.querySelector('#progress_banner > p')
  

  function updateFilm(evt) {
    if (evt.lengthComputable) {
      var percentLoaded = Math.round((evt.loaded / evt.total) * 100);
      if (percentLoaded < 100) {
        pFilm.style.transform = `scaleX(${percentLoaded / 100})`;
        tFilm.textContent = percentLoaded + '%';
      }
    }
  }

  function handleFilm(evt) {
    reader = new FileReader();
    reader.onprogress = updateFilm;
    reader.onloadstart = function(e) {
      pFilm.style.transform = 'scaleX(0)';
      tFilm.textContent = '0%';      
    };
    reader.onload = function(e) {
      pFilm.style.transform = 'scaleX(1)';
      tFilm.textContent = '100%';
    }
    reader.readAsBinaryString(evt.target.files[0]);
  }


  function updatePoster(evt) {
    if (evt.lengthComputable) {
      var percentLoaded = Math.round((evt.loaded / evt.total) * 100);
      if (percentLoaded < 100) {
        pPoster.style.transform = `scaleX(${percentLoaded / 100})`;
        tPoster.textContent = percentLoaded + '%';
      }
    }
  }
  function handlePoster(evt) {
    reader = new FileReader();
    reader.onprogress = updatePoster;
    reader.onloadstart = function(e) {
      pPoster.style.transform = 'scaleX(0)';
      tPoster.textContent = '0%';      
    };
    reader.onload = function(e) {
      pPoster.style.transform = 'scaleX(1)';
      tPoster.textContent = '100%';
    }
    reader.readAsBinaryString(evt.target.files[0]);
  }


  function updateBanner(evt) {
    if (evt.lengthComputable) {
      var percentLoaded = Math.round((evt.loaded / evt.total) * 100);
      if (percentLoaded < 100) {
        pBanner.style.transform = `scaleX(${percentLoaded / 100})`;
        tBanner.textContent = percentLoaded + '%';
      }
    }
  }
  function handleBanner(evt) {
    reader = new FileReader();
    reader.onprogress = updateBanner;
    reader.onloadstart = function(e) {
      pBanner.style.transform = 'scaleX(0)';
      tBanner.textContent = '0%';      
    };
    reader.onload = function(e) {
      pBanner.style.transform = 'scaleX(1)';
      tBanner.textContent = '100%';
    }
    reader.readAsBinaryString(evt.target.files[0]);
  }

  document.querySelector('#film__file').addEventListener('change', handleFilm, false);
  document.querySelector('#poster__file').addEventListener('change', handlePoster, false);
  document.querySelector('#banner__file').addEventListener('change', handleBanner, false);

  

  




  