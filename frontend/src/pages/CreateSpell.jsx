import CreateSpell from "../assets/createspellpage.jpg";

const CreateSpells = () => {
  return (
    <div className="create-source-container grid grid-cols-12 h-screen">
      <div className="col-span-4 h-full">
        <img
          className="w-full object-cover h-full"
          src={CreateSpell}
          alt="Charging a random spell"
        />
      </div>
    </div>
  );
};

export default CreateSpells;
