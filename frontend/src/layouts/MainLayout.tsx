import React from 'react';
import { PictureSlideShow } from '../components/medievalMemes';
import { Grid } from '@mui/material';
import { Bestlist } from '../components/bestlist';
import { Newsfeed } from '../components/newsfeed';
import { WebSocketProvider } from '../contexts/newsContext';

export const MainLayout: React.FC = () => {
  return (
    <WebSocketProvider>
      <div
        style={{
          width: '100%',
          minHeight: '100vh',
          // background: 'url("/images/background/background.jpg")',
          background: 'black',
        }}
      >
        <Grid
          container
          spacing={1}
          style={{
            width: '100%',
            minHeight: '100vh',
            // background: 'url("/images/background/background.jpg")',
            background: 'black',
          }}
        >
          <Grid item xs={6}>
            <Bestlist withAutoScroll={false} />
          </Grid>
          <Grid item xs={6}>
            <Grid container spacing={1}>
              <Grid item xs={6}>
                <a href="/game">
                  <img
                    className="animate__animated animate__pulse"
                    src="images/logo.png"
                    style={{
                      width: '100%',
                      animationIterationCount: 'infinite',
                    }}
                  />
                </a>
              </Grid>
              <Grid item xs={6}>
                <PictureSlideShow />
              </Grid>
              <Grid item xs={12}>
                <Newsfeed animation="animate__animated animate__fadeInUp" />
              </Grid>
            </Grid>
          </Grid>
        </Grid>
      </div>
    </WebSocketProvider>
  );
};
