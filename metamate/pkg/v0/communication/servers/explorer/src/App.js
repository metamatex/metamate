import React, { useState, useEffect } from 'react';
import './App.css';
import GraphiQLExplorer from "graphiql-explorer"
import GraphiQL from "graphiql"
import "graphiql/graphiql.css"
import { getIntrospectionQuery, buildClientSchema } from "graphql"

const parameters = {};

window.location.search
    .substr(1)
    .split(`&`)
    .forEach(function(entry) {
        let eq = entry.indexOf(`=`);

        if (eq >= 0) {
            parameters[decodeURIComponent(entry.slice(0, eq))] = decodeURIComponent(
                entry.slice(eq + 1)
            )
        }
    });

function updateURL() {
    window.history.replaceState(null, null, locationQuery(parameters))
}

function locationQuery(params) {
    return (
        `?` +
        Object.keys(params)
            .map(function(key) {
                return encodeURIComponent(key) + `=` + encodeURIComponent(params[key])
            })
            .join(`&`)
    )
}

function onEditVariables(newVariables) {
    parameters.variables = newVariables;
    updateURL()
}
function onEditOperationName(newOperationName) {
    parameters.operationName = newOperationName;
    updateURL()
}

function graphQLFetcher(graphQLParams) {
    return fetch(window.graphqlEndpoint, {
        method: `post`,
        headers: {
            Accept: `application/json`,
            "Content-Type": `application/json`,
        },
        body: JSON.stringify(graphQLParams),
    }).then(function (response) {
        return response.json()
    })
}

const storedExplorerPaneState =
    typeof parameters.explorerIsOpen !== `undefined`
        ? parameters.explorerIsOpen !== `false`
        : window.localStorage
        ? window.localStorage.getItem(`graphiql:graphiqlExplorerOpen`) !== `false`
        : false;

const storeIsDarkState =
    typeof parameters.isDark !== `undefined`
        ? parameters.isDark !== `false`
        : window.localStorage
        ? window.localStorage.getItem(`graphiql:isDark`) !== `false`
        : false;

let mounted = false;

function App() {
    const [schema, setSchema] = useState(null);
    const [query, setQuery] = useState(null);
    const [explorerIsOpen, setExplorerIsOpen] = useState(storedExplorerPaneState);
    const [isDark, setIsDark] = useState(storeIsDarkState);
    const [graphiql, setGraphiql] = useState(null);

    useEffect(() => {
        graphQLFetcher({
            query: getIntrospectionQuery(),
        }).then(result => {
            setSchema(buildClientSchema(result.data));

            mounted = true;
        });

        setQuery(parameters.query || (window.localStorage && window.localStorage.getItem(`graphiql:query`)) || window.defaultQuery);
    }, [mounted]);

    const handleEditQuery = query => {
        parameters.query = query;
        updateURL();
        setQuery(query);
    };

    const handleToggleExplorer = query => {
        const newExplorerIsOpen = !explorerIsOpen;
        if (window.localStorage) {
            window.localStorage.setItem(
                `graphiql:graphiqlExplorerOpen`,
                newExplorerIsOpen
            )
        }

        parameters.explorerIsOpen = newExplorerIsOpen;
        updateURL();
        setExplorerIsOpen(newExplorerIsOpen);
    };

    const handleToggleDark = query => {
        const newIsDark = !isDark;
        if (window.localStorage) {
            window.localStorage.setItem(
                `graphiql:isDark`,
                newIsDark
            )
        }

        parameters.isDark = newIsDark;
        updateURL();
        setIsDark(newIsDark);
    };

    return (
    <div className={"graphiql-container" + (isDark ? ' isDark' : '') + (parameters.isCompact ? ' isCompact' : '')}>
        {parameters.color ? <style dangerouslySetInnerHTML={{__html: `#root .execute-button { border: none; background: ` + parameters.color + ` }`}}></style>: null}
            <GraphiQLExplorer
                schema={schema}
                query={query}
                onEdit={handleEditQuery}
                explorerIsOpen={explorerIsOpen}
                onToggleExplorer={handleToggleExplorer}
                onRunOperation={operationName =>
                    graphiql.handleRunQuery(operationName)
                }
                colors={{
                    keyword: '#ff7d00',
                    // OperationName, FragmentName
                    def: '#ff0025',
                    // FieldName
                    property: '#008cff',
                    // FieldAlias
                    qualifier: '#1C92A9',
                    // ArgumentName and ObjectFieldName
                    attribute: '#ed06ff',
                    number: '#2882F9',
                    string: '#35ff00',
                    // Boolean
                    builtin: '#D47509',
                    // Enum
                    string2: '#0B7FC7',
                    variable: '#397D13',
                    // Type
                    atom: '#CA9800',
                }}
            />
            <GraphiQL
                ref={setGraphiql}
                fetcher={graphQLFetcher}
                schema={schema}
                query={query}
                onEditQuery={handleEditQuery}
                onEditVariables={onEditVariables}
                onEditOperationName={onEditOperationName}
            >
                    {parameters.isCompact ? null : <GraphiQL.Logo>
                    <img className="logo" src="/static/logo/blue_transparent.svg" alt=""/>
        </GraphiQL.Logo>}
                <GraphiQL.Toolbar>
        {parameters.isCompact ? null : <GraphiQL.Button
                        label="Prettify"
                        title="Prettify Query (Shift-Ctrl-P)"
                        onClick={() => graphiql.handlePrettifyQuery()}
                    />}
                    {parameters.isCompact ? null : <GraphiQL.Button
                        label="History"
                        title="Show History"
                        onClick={() => graphiql.handleToggleHistory()}
                    />}
                    {parameters.isCompact ? null : <GraphiQL.Button
                        label="Explorer"
                        title="Toggle Explorer"
                        onClick={handleToggleExplorer}
                    />}
                    {parameters.isCompact ? null : <GraphiQL.Button
                        label={isDark ? '🌕 Light' : '🌑 Dark'}
                        title="Toggle Dark Mode"
                        onClick={handleToggleDark}
                    />}
                    {parameters.isCompact ? <a style={{color: parameters.color}} href={"http://" + window.location.host + window.location.pathname + "?query=" +encodeURI(parameters.query)} target="_blank" rel="noopener noreferrer" className={"toolbar-button open-in-explorer-button"}>Open in explorer</a> : null}
                </GraphiQL.Toolbar>
            </GraphiQL>
        </div>
    );
}

export default App;
