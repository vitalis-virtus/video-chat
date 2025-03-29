import React from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";

const CreateRoom = () => {
  const navigate = useNavigate();

  const create = async (e) => {
    e.preventDefault();

    try {
      const response = await axios.post("http://localhost:9000/channels");
      const { id } = response.data;

      navigate(`/channels/${id}`, { state: { id: id } });
    } catch (error) {
      console.error("Failed to create room:", error);
    }
  };

  return (
    <div>
      <button onClick={create}>Create Room</button>
    </div>
  );
};

export default CreateRoom;
