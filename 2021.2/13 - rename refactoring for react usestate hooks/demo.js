import React, {useState} from 'react';

function Demo() {
    const [visitors, setVisitors] = useState(0);

    return (<div>
        <p>{visitors} visitors have seen this page.</p>
        <button onClick={() => setVisitors(visitors + 1)}>Visit!</button>
    </div>);
}