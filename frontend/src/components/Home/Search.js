import React from "react";
import agent from "../../agent";

const Search = (props) => {
    const handleChange = event => {
        if(event.target.value.length >= 3){
            let title = event.target.value;
            props.onSearch(
                title,
                (page) => agent.Items.byTitle(title, page),
                agent.Items.byTitle(title)
            );
        } else {
            props.onSearch(
                "",
                (page) => agent.Items.byTitle("", page),
                agent.Items.byTitle("")
            );
        }
    };

    return (
        <span>
          <input id="search-box" name="search-box" type='text' placeholder='What is it that you truly desire?' onChange={handleChange} />
        </span>
    );
};

export default Search;
