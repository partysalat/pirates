import React, { useEffect, useState } from 'react';
import { SwitchTransition, CSSTransition } from 'react-transition-group';
import './style.css';
const duration = 500;

const MEMES: string[] = [
  'blaubaer.webp',
  'legurk.webp',
  'gisela.webp',
  'piraten_gemuese.webp',
  'nerf.webp',
  'flunkyball.webp',
  'schlackerschwerter.webp',
];

export const PictureSlideShow = () => {
  const [image, setImage] = useState(MEMES[0]);
  useEffect(() => {
    const interval = setInterval(() => {
      setImage(MEMES[Math.floor(Math.random() * MEMES.length)]);
    }, 15000);
    return () => {
      clearInterval(interval);
    };
  }, []);
  // return <div style={{ width: '100%', height: "100%" }} src={`images/memes/${image}`} />;

  return (
    <SwitchTransition>
      <CSSTransition key={image} timeout={duration} classNames="fade">
        <div
          style={{
            height: '100%',
            backgroundImage: `url(images/slideshow/${image})`,
            backgroundSize: 'contain',
            backgroundRepeat: 'no-repeat',
            backgroundPosition: 'center',
          }}
        />
      </CSSTransition>
    </SwitchTransition>
  );
};
