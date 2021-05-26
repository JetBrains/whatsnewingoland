import {useState} from "react";
import clsx from "clsx";

function Demo() {
    const [visitors, setVisitors] = useState(0);

    return (<div>
        <p>{visitors} visitors have seen this page.</p>
        <button
            className={clsx('bt')}
            onClick={() => setVisitors(visitors + 1)}
        >Visit!</button>
    </div>);
}