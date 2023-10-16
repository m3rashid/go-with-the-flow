import React from 'react';

import Message from './components/message';

function App() {
  return (
    <div>
      Hello
      {Array(100)
        .fill(4)
        .map((t) => (
          <div>{t}</div>
        ))}
      <Message
        hideUndo={false}
        onUndo={() => console.log('undo')}
        hideDismiss={false}
        onDismiss={() => console.log('dismiss')}
        message='Hello'
        description='This is a hello message'
      />
    </div>
  );
}

export default App;
