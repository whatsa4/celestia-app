package appconsts

// MalleatedTxBytes is the overhead bytes added to a normal transaction after
// malleating it. 32 for the original hash, 4 for the uint32 share_index, and 3
// for protobuf
const MalleatedTxBytes = 32 + 4 + 3