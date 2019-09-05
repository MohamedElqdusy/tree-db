/* tree nodes tables */
CREATE TABLE nodes(id INTEGER NOT NULL PRIMARY KEY, content text NOT NULL, tree_root INTEGER NOT NULL,parent INTEGER NULL REFERENCES nodes(id),path TEXT NOT NULL);

/* node's height based on path splits */
CREATE OR REPLACE  FUNCTION height(node_id integer) RETURNS integer AS $$
declare node_path text;
declare node_height integer;

BEGIN
SELECT path into node_path from nodes where id = node_id;
select array_length( string_to_array(node_path, ','),1 ) into node_height;
RETURN node_height -2; /* -2 discarding the root and the current node from path splits count */
END; $$
LANGUAGE PLPGSQL;

/* returns rows of the node and its childs rows from the node's path */
CREATE OR REPLACE  FUNCTION get_node_with_childs(node_id integer) RETURNS SETOF nodes AS $$
declare node_path text;
BEGIN
SELECT path into node_path from nodes where id = node_id;
RETURN QUERY select * from nodes where path like concat(node_path,'%') ORDER BY height(id);
END; $$
LANGUAGE PLPGSQL;


CREATE OR REPLACE  FUNCTION change_Parent(given_node_id integer, new_parent integer) RETURNS void AS $$
DECLARE row record;
BEGIN
/* 1- updtae the parent of the given node */
UPDATE nodes SET parent = new_parent where id = given_node_id; 
/* 2- update the path of the given node and its childs */ 
    FOR row in SELECT id FROM get_node_with_childs(given_node_id) LOOP
        perform change_node_Path_with_parent_path(row.id);
    END LOOP;    
END; $$
LANGUAGE PLPGSQL;

/* change node's path with its current parent's path */
CREATE OR REPLACE  FUNCTION change_node_Path_with_parent_path(given_node_id integer) RETURNS void AS $$
declare parent_path text;
declare parent_id integer;
BEGIN 
select parent into parent_id from nodes where id = given_node_id;
select path into parent_path from nodes where id = parent_id;
/* update the path of the given node with its parent's path */ 
UPDATE nodes SET path = concat(parent_path, id,',') where id = given_node_id;
END; $$
LANGUAGE PLPGSQL;
